package xetcd

import (
	"context"
	"fmt"
	"lark/com/pkgs/xlog"
	"net"
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

/*
处理 新的微服务注册到etcd的流程
*/

const (
	TIME_TO_LIVE = 10 //TTL 生存时间 秒
)

type RegEtcd struct {
	cli          *clientv3.Client
	endpoints    []string
	serviceValue string
	serviceKey   string
	ttl          int
	host         string
	port         int
	schema       string
	serviceName  string
	//接收重新注册标识
	reChan chan struct{}
	//判断当前是否注册成功
	connected bool
	//接收过期回调
	kresp  <-chan *clientv3.LeaseKeepAliveResponse
	cancel context.CancelFunc
}

//拼接 etcd的key的前缀
// "项目标识:///服务名/"
func GetPrefix(schema, servername string) string {
	return fmt.Sprintf("%s:///%s/", schema, servername)
}

/*
将微服务注册进etcd
schema 项目标识
endpoints etcd的服务地址列表
myHost 要注册的微服务的ip地址
myPort 要注册的微服务的端口
serviceName 要注册的微服务名
ttl 有效时间长度(秒)
*/
func RegisterEtcd(schema string, endpoints []string, myHost string, myPort int, serviceName string, ttl int) {
	serviceValue := net.JoinHostPort(myHost, strconv.Itoa(myPort))
	serviceKey := GetPrefix(schema, serviceName) + serviceValue
	rEtcd := &RegEtcd{
		endpoints:    endpoints,
		serviceValue: serviceValue,
		serviceKey:   serviceKey,
		ttl:          ttl,
		host:         myHost,
		port:         myPort,
		schema:       schema,
		serviceName:  serviceName,
		reChan:       make(chan struct{}),
	}
	rEtcd.Register()
	rEtcd.reRegister()
}

//将service 注册进etcd
func (r *RegEtcd) Register() (err error) {
	defer func() {
		if err != nil {
			xlog.Error(err.Error())
			r.reChan <- struct{}{}
		}
	}()

	var cli *clientv3.Client
	//创建etcd客户端
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return
	}
	//创建etcd lease
	ctx, cancel := context.WithCancel(context.Background())
	var resp *clientv3.LeaseGrantResponse
	resp, err = cli.Grant(ctx, int64(r.ttl))
	if err != nil {
		return
	}
	//向etcd写入数据
	if _, err = cli.Put(ctx, r.serviceKey, r.serviceValue, clientv3.WithLease(resp.ID)); err != nil {
		return
	}
	//获取过期回调通道
	var kresp <-chan *clientv3.LeaseKeepAliveResponse
	kresp, err = cli.KeepAlive(ctx, resp.ID)
	if err != nil {
		return
	}
	r.cli = cli
	r.connected = true
	r.kresp = kresp
	r.cancel = cancel
	return
}

//重新注册流程 包括过期 和 注册失败
func (r *RegEtcd) reRegister() {
	go func() {
		var ok bool
		for {
			if r.connected {
				//接收过期回调
				_, ok = <-r.kresp
				r.connected = ok
				if !ok {
					xlog.Error("租约失效")
					r.cancel()
					r.reChan <- struct{}{}
				}
			} else {
				//处理重新注册
				_, ok = <-r.reChan
				if !ok {
					return
				}
				time.Sleep(1 * time.Second)
				r.Register()
			}
		}
	}()
}

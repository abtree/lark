package xetcd

import (
	"context"
	"lark/com/pkgs/xlog"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

/*
处理从etcd中查询服务
*/

//Resolver 需要实现 resolver.Builder
type GetEtcd struct {
	prefix             string
	cli                *clientv3.Client
	watchStartRevision int64
}

//创建获取etcd对象
func NewGetEtcd(endpoints []string, schema, serviceName string, userName, password string) *GetEtcd {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		Username:  userName,
		Password:  password,
	})
	if err != nil {
		xlog.Error(err.Error())
		return nil
	}
	return &GetEtcd{
		prefix: GetPrefix(schema, serviceName),
		cli:    cli,
	}
}

/* 获取服务列表
timeout 超时时间
update 如果传入update回调 则会监听数据更新
*/
func (r *GetEtcd) Get(timeout time.Duration, update func(*clientv3.Event)) []*mvccpb.KeyValue {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	resp, err := r.cli.Get(ctx, r.prefix, clientv3.WithPrefix())
	if err != nil {
		return nil
	}
	r.watchStartRevision = resp.Header.Revision + 1
	//如果需要监控更新
	if update != nil {
		go r.watch(update)
	}
	return resp.Kvs
}

//监控更新
func (r *GetEtcd) watch(update func(*clientv3.Event)) {
	rch := r.cli.Watch(context.Background(), r.prefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			update(ev)
		}
	}
}

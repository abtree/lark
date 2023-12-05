package xgrpc

import (
	"context"
	"fmt"
	"io"
	"lark/com/pkgs/xetcd"
	"lark/com/pkgs/xlog"
	"lark/pb"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

const (
	CONST_DURATION_GRPC_TIMEOUT_SECOND = 5 * time.Second
)

/*
需要实现接口：resolver.Builder resolver.Resolver
*/
type Resolver struct {
	srv            *pb.CfgServers
	etcdcfg        *pb.Jetcd
	cc             resolver.ClientConn
	grpcClientConn *grpc.ClientConn
	closer         io.Closer
	etcd           *xetcd.GetEtcd
	addrList       map[string]string
}

var (
	resolvers     = map[string]*Resolver{}
	resolverMutex sync.RWMutex
)

// resolver.Builder 需要实现的接口
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc
	back := r.etcd.Get(CONST_DURATION_GRPC_TIMEOUT_SECOND, r.watch)
	if back != nil {
		r.addrList = map[string]string{}
		arr := []resolver.Address{}
		for i := range back {
			r.addrList[string(back[i].Key)] = string(back[i].Value)
			arr = append(arr, resolver.Address{Addr: string(back[i].Value)})
		}
		r.cc.UpdateState(resolver.State{Addresses: arr})
	}
	return r, nil
}

// 监听变化
func (r *Resolver) watch(ev *clientv3.Event) {
	key := string(ev.Kv.Key)
	switch ev.Type {
	case mvccpb.PUT:
		if _, ok := r.addrList[key]; !ok {
			r.addrList[key] = string(ev.Kv.Value)
			r.updateAddress()
		}
	case mvccpb.DELETE:
		if _, ok := r.addrList[key]; ok {
			delete(r.addrList, key)
			r.updateAddress()
		}
	}
}

func (r *Resolver) updateAddress() {
	arr := []resolver.Address{}
	for _, v := range r.addrList {
		arr = append(arr, resolver.Address{Addr: v})
	}
	r.cc.UpdateState(resolver.State{Addresses: arr})
}

func (r *Resolver) Scheme() string {
	return r.etcdcfg.Schema
}

// resolver.Resolver 需要实现的接口
func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {

}

func (r *Resolver) Close() {

}

// 新建Resolver
func NewResolver(cfg *pb.Jetcd, srv *pb.CfgServers) *Resolver {
	r := &Resolver{
		srv:     srv,
		etcdcfg: cfg,
		etcd:    xetcd.NewGetEtcd(cfg.Endpoints, cfg.Schema, srv.Name, cfg.Username, cfg.Password),
	}
	resolverMutex.Lock()
	defer resolverMutex.Unlock()
	resolver.Register(r)
	r.grpcClientConn, _ = r.newGrpcClientConn()
	return r
}

func (r *Resolver) newGrpcClientConn() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	// 开启链路追踪
	// if r.opt.Tracing.Enabled == true && r.opt.Tracing.Tracer != nil {
	// 	opts = append(opts, grpc.WithBlock())
	// 	opts = append(opts, grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(r.opt.Tracing.Tracer)))
	// }
	if r.srv.Cert {
		//开启认证
		creds, err := credentials.NewClientTLSFromFile(r.srv.CertPem, r.srv.CertName)
		if err != nil {
			xlog.Error(err.Error())
		} else {
			opts = append(opts, grpc.WithTransportCredentials(creds))
		}
	}
	ctx, _ := context.WithTimeout(context.Background(), CONST_DURATION_GRPC_TIMEOUT_SECOND)
	conn, err := grpc.DialContext(ctx, xetcd.GetPrefix(r.etcdcfg.Schema, r.srv.Name), opts...)
	if err != nil {
		xlog.Error(err.Error())
		return nil, err
	}
	return conn, nil
}

package xgrpc

import (
	"io"
	"lark/pb"

	"google.golang.org/grpc"
)

type ClientDialOption struct {
	ServiceName string
	etcd        *pb.Jetcd
	srv         *pb.CfgServers
	closer      io.Closer
}

func NewClientDialOption(etcd *pb.Jetcd, server *pb.CfgServers, jaeger *pb.Jjaeger, clientName string) *ClientDialOption {
	ret := &ClientDialOption{
		ServiceName: server.Name,
		etcd:        etcd,
		srv:         server,
	}
	if jaeger.Enable {
		//链路追踪
		//tracer, closer, _ = xtracer.NewTracer(clientName, jaeger)
		// opt.Tracing = &conf.Tracing{Tracer: tracer, Enabled: jaeger.Enabled}
		// ret.closer = closer
	}
	return ret
}

func (opt *ClientDialOption) GetClientConn() *grpc.ClientConn {
	key := opt.etcd.Schema + opt.ServiceName
	resolverMutex.RLock()
	r, ok := resolvers[key]
	resolverMutex.RUnlock()
	if ok {
		return r.grpcClientConn
	}
	r = NewResolver(opt.etcd, opt.srv)
	resolverMutex.Lock()
	defer resolverMutex.Unlock()
	resolvers[key] = r
	return r.grpcClientConn
}

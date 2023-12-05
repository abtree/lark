package xgrpc

import (
	"io"
	"lark/com/pkgs/xetcd"
	"lark/com/pkgs/xlog"
	"lark/pb"
	"net"
	"strconv"
	"time"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	grpc   *pb.Jgrpc
	etcd   *pb.Jetcd
	server *pb.CfgServers
	jaeger *pb.Jjaeger
}

func RecoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	})
}

// 启动grpc server第一步 : 创建GrpcServer
func NewGrpcServer(grpc *pb.Jgrpc, etcd *pb.Jetcd, srv *pb.CfgServers, jaeger *pb.Jjaeger) *GrpcServer {
	return &GrpcServer{
		grpc:   grpc,
		etcd:   etcd,
		server: srv,
		jaeger: jaeger,
	}
}

// 启动grpc server第二步 : 创建grpc.Server
func (s *GrpcServer) NewServer() (srv *grpc.Server, closer io.Closer) {
	opts := []grpc.ServerOption{}
	opts = append(opts, grpc.UnaryInterceptor(grpc_recovery.UnaryServerInterceptor()))
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(s.grpc.MaxConnectionIdle) * time.Millisecond,
		MaxConnectionAge:      time.Duration(s.grpc.MaxConnectionAge) * time.Millisecond,
		MaxConnectionAgeGrace: time.Duration(s.grpc.MaxConnectionAgeGrace) * time.Millisecond,
		Time:                  time.Duration(s.grpc.Time) * time.Millisecond,
		Timeout:               time.Duration(s.grpc.Timeout) * time.Millisecond,
	})
	opts = append(opts, keepParams)
	if s.server.Cert {
		//tls 认证
		creds, err := credentials.NewServerTLSFromFile(s.server.CertPem, s.server.CertKey)
		if err != nil {
			xlog.Error(err.Error())
		} else {
			opts = append(opts, grpc.Creds(creds))
		}
	}
	if s.jaeger.Enable {
		//todo 链路追踪
	}
	if s.grpc.StreamsLimit > 0 {
		// 一个连接中最大并发Stream数
		opts = append(opts, grpc.MaxConcurrentStreams(uint32(s.grpc.StreamsLimit)))
	}
	if s.grpc.MaxRecvMsgSize > 0 {
		// 允许接收的最大消息长度
		opts = append(opts, grpc.MaxRecvMsgSize(int(s.grpc.MaxRecvMsgSize)))
	}
	srv = grpc.NewServer(opts...)
	return
}

// 启动grpc server第三步 : 启动server
func (s *GrpcServer) RunServer(server *grpc.Server) {
	defer server.GracefulStop()

	address := "0.0.0.0:" + strconv.Itoa(int(s.server.Port))
	listener, err := net.Listen("tcp", address)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	if s.grpc.ConnectionLimit > 0 {
		// 最大并发连接数
		listener = netutil.LimitListener(listener, int(s.grpc.ConnectionLimit))
	}
	//注册服务
	xetcd.RegisterEtcd(s.etcd.Schema, s.etcd.Endpoints, s.server.Host, int(s.server.Port), s.server.Name, xetcd.TIME_TO_LIVE)
	//启动监听
	xlog.Info("grpc server start at ", address)
	err = server.Serve(listener)
	if err != nil {
		xlog.Error("grpc server error start at", address, err.Error())
	}
}

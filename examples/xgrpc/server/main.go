package main

import (
	"context"
	"lark/com/pkgs/xgrpc"
	"lark/com/pkgs/xloadcfg"
	"lark/com/pkgs/xlog"
	"lark/pb"
)

func main() {
	cfg := &pb.MsgSysConfigs{}
	xloadcfg.Run("../../../configs", cfg, nil)
	xlog.Shared(cfg.Logger, "examples")

	svrcfg := cfg.Servers[2] //lark_config
	svrcfg.Host = "127.0.0.1"

	server := xgrpc.NewGrpcServer(cfg.Grpc, cfg.Etcd, svrcfg, cfg.Jaeger)
	srv, closer := server.NewServer()
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()
	pb.RegisterSrvServiceServer(srv, &SvcConfigServer{})
	server.RunServer(srv)
}

type SvcConfigServer struct {
	pb.UnimplementedSrvServiceServer
}

func (s *SvcConfigServer) SrvSvc(ctx context.Context, req *pb.SvcReq) (*pb.SvcResp, error) {
	xlog.Info(req)
	return &pb.SvcResp{
		Code: 0,
		Msg:  "Ok",
	}, nil
}

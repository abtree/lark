package api

import (
	"context"
	"lark/com/pkgs/xgrpc"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"lark/pb"

	"google.golang.org/grpc/metadata"
)

type IGrpcServer interface {
	pb.SrvServiceServer
	MainInstance
}

type GrpcServer struct {
	pb.UnimplementedSrvServiceServer
	grpcServer *xgrpc.GrpcServer
	MsgServer
}

func (s *GrpcServer) Init() error {
	return nil
}

func (s *GrpcServer) RunLoop() {
	cfg := App.SysCfg
	s.grpcServer = xgrpc.NewGrpcServer(cfg.Grpc, cfg.Etcd, App.SrvCfg, cfg.Jaeger)
	srv, closer := s.grpcServer.NewServer()
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()
	pb.RegisterSrvServiceServer(srv, s)
	s.grpcServer.RunServer(srv)
}

func (s *GrpcServer) Destory() {

}

func (s *GrpcServer) SrvSvc(ctx context.Context, req *pb.SvcReq) (*pb.SvcResp, error) {
	typ := uint32(0)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		xlog.Warn("grpc missing metadata")
	}
	if val, ok := md[ServerType]; ok {
		typ = utils.StrToUint32(val[0])
	}
	return s.Handle(typ, req)
}

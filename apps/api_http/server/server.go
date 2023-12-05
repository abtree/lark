package server

import (
	"lark/apps/api_http/router"
	"lark/com/api"
	"lark/com/pkgs/xgin"
	"lark/pb"
)

type ApiServer interface {
	api.MainInstance
}

type apiServer struct {
	ginServer  *xgin.GinServer
	grpcServer *api.GrpcServer
}

func NewApiServer() api.MainInstance {
	return &apiServer{
		ginServer:  xgin.NewServer(),
		grpcServer: &api.GrpcServer{},
	}
}

func (s *apiServer) Init() error {
	api.NewGrpcService(pb.ServerType_Auth)
	api.NewGrpcService(pb.ServerType_Giftpack)

	router.Register(s.ginServer.Engine)
	return s.grpcServer.Init()
}

func (s *apiServer) RunLoop() {
	go s.ginServer.Run(int(api.App.SrvCfg.HttpPort))
	s.grpcServer.RunLoop()
}

func (s *apiServer) Destory() {
	s.grpcServer.Destory()
}

// 注册消息处理函数
func (s *apiServer) RegisterMsg() {
	s.RegisterCommonMsg()
}

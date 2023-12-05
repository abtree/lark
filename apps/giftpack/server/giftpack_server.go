package server

import (
	"lark/apps/giftpack/service"
	"lark/com/api"
	"lark/pb"
)

type GiftpackServer interface {
	api.IGrpcServer
}

type giftpackServer struct {
	api.GrpcServer
	gpService *service.GiftpackService
}

func NewGiftpackServer(svc *service.GiftpackService) GiftpackServer {
	return &giftpackServer{
		gpService: svc,
	}
}

func (s *giftpackServer) Init() error {
	api.NewGrpcService(pb.ServerType_Config)

	s.gpService.Init()
	return nil
}

func (s *giftpackServer) Destory() {
	s.gpService.Exit()
}

// 注册消息处理函数
func (s *giftpackServer) RegisterMsg() {
	s.RegisterCommonMsg()
}

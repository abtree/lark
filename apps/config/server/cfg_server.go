package server

import (
	"lark/apps/config/service"
	"lark/com/api"
)

type ConfigServer interface {
	api.IGrpcServer
}

type configServer struct {
	api.GrpcServer
	cfgService *service.CfgService
}

func NewConfigServer(svc *service.CfgService) ConfigServer {
	return &configServer{
		cfgService: svc,
	}
}

func (s *configServer) Init() error {
	s.cfgService.LoadAllConfigs()

	return nil
}

// 注册消息处理函数
func (s *configServer) RegisterMsg() {
	s.RegisterCommonMsg()
}

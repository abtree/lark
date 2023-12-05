package server

import (
	"lark/apps/api_ws/websocket"
	"lark/com/api"
	"lark/com/pkgs/xgin"
	"lark/com/tools"
	"lark/pb"
)

type WSServer interface {
	api.MainInstance
}

type wsServer struct {
	ginServer *xgin.GinServer
	hub       *websocket.Hub //暂时一个server只有一个hub（可以改为多个模式）
	cfg       *pb.MsgAllConfigs
}

func NewWSServer() api.MainInstance {
	return &wsServer{
		ginServer: xgin.NewServer(),
	}
}

func (s *wsServer) Init() (err error) {
	api.NewGrpcService(pb.ServerType_Config)
	s.cfg, err = tools.GetAllCfg()
	if err != nil {
		return
	}
	s.hub = websocket.NewHub(s.cfg.Configs.Websocket, s.OnProto)
	s.Router()
	return nil
}

func (s *wsServer) RunLoop() {
	s.hub.Run()
	s.ginServer.Run(int(api.App.SrvCfg.HttpPort))
}

func (s *wsServer) Destory() {
	s.hub.Exit()
}

func (s *wsServer) Router() {
	s.ginServer.Engine.GET("/", s.hub.Handler)
}

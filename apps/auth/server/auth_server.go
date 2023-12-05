package server

import (
	"lark/apps/auth/service"
	"lark/com/api"
	"lark/com/pkgs/xmysql"
	"lark/com/pkgs/xredis"
	"lark/pb"
)

type AuthServer interface {
	api.IGrpcServer
}

type authServer struct {
	//集成grpc服务
	api.GrpcServer
	//逻辑处理service
	authService *service.AuthService
}

func NewAuthServer(auth *service.AuthService) AuthServer {
	return &authServer{
		authService: auth,
	}
}

func (s *authServer) Init() error {
	xmysql.NewMysqlClient(api.App.SysCfg.Mysql)
	xredis.NewRedisClient(api.App.SysCfg.Redis)

	api.NewGrpcService(pb.ServerType_GetId)
	return nil
}

// 注册消息处理函数
func (s *authServer) RegisterMsg() {
	s.RegisterCommonMsg()
}

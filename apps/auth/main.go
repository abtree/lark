package main

import (
	"lark/apps/auth/server"
	"lark/apps/auth/service"
	"lark/com/api"
	"lark/pb"
)

/*
	用于登录验证的微服务
*/

func main() {
	svc := service.NewAuthService()
	srv := server.NewAuthServer(svc)
	api.Run(pb.ServerType_Auth, srv)
}

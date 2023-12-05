package main

import (
	"lark/apps/config/server"
	"lark/apps/config/service"
	"lark/com/api"
	"lark/pb"
)

/* 处理配置文件的服务
 */
func main() {
	svc := service.NewCfgService()
	srv := server.NewConfigServer(svc)
	api.Run(pb.ServerType_Config, srv)
}

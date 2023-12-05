package main

import (
	"lark/apps/getid/server"
	"lark/apps/getid/service"
	"lark/com/api"
	"lark/pb"
)

func main() {
	svc := service.NewGetidService()
	srv := server.NewGetidServer(svc)
	api.Run(pb.ServerType_GetId, srv)
}

package main

import (
	"lark/apps/giftpack/server"
	"lark/apps/giftpack/service"
	"lark/com/api"
	"lark/pb"
)

func main() {
	svc := service.NewGiftpackService()
	srv := server.NewGiftpackServer(svc)
	api.Run(pb.ServerType_Giftpack, srv)
}

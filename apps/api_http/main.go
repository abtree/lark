package main

import (
	"lark/apps/api_http/server"
	"lark/com/api"
	"lark/pb"
)

func main() {
	srv := server.NewApiServer()
	api.Run(pb.ServerType_ApiHttp, srv)
}

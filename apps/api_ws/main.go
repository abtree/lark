package main

import (
	"lark/apps/api_ws/server"
	"lark/com/api"
	"lark/pb"
)

func main() {
	srv := server.NewWSServer()
	api.Run(pb.ServerType_WebSocket, srv)
}

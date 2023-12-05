package server

import (
	"lark/apps/api_ws/websocket"
	"lark/com/pkgs/xlog"
)

func (s *wsServer) OnProto(task *websocket.Wstask) {
	xlog.Debug(task)
	s.hub.Broadcast(task.Msg)
}

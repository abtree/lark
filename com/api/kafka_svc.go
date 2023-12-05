package api

import (
	"lark/com/pkgs/xkafka"
	"lark/pb"

	"google.golang.org/protobuf/proto"
)

func NewKafkaService() {
	xkafka.NewWriter(App.SysCfg.Kafka)
}

// 通过kafka向特定server发送消息
func SendToKafka(svrType pb.ServerType, serverId uint32, msgid uint32, msg proto.Message) error {
	return xkafka.SendToServer(App.GetSvrNameBySvrType(svrType), serverId, msgid, msg)
}

// 通过kafka向特定server广播消息消息
func BroadcastToKafka(svrType pb.ServerType, msgid uint32, msg proto.Message) error {
	return xkafka.BroadcastToServer(App.GetSvrNameBySvrType(svrType), msgid, msg)
}

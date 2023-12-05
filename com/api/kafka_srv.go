package api

import (
	"lark/com/pkgs/xkafka"
	"lark/com/pkgs/xlog"
	"lark/com/utils"

	"google.golang.org/protobuf/proto"
)

type KafkaServer struct {
	MsgServer
}

func (s *KafkaServer) Init() error {
	xkafka.NewReader(App.SrvCfg.Name, App.SrvCfg.ServerId, App.SysCfg.Kafka)
	return nil
}

func (s *KafkaServer) RunLoop() {
	xkafka.RunReader(s.Handle)
}

func (s *KafkaServer) Destory() {
	xkafka.ExitReader()
}

func (s *KafkaServer) Handle(msgid uint32, data []byte) {
	if s.handles == nil {
		xlog.Error(utils.CONST_kafka_NOT_HANDLE)
		return
	}
	hand := s.getHandler(0, msgid)
	if hand == nil {
		xlog.Error(utils.CONST_kafka_NOT_HANDLE)
		return
	}
	//解析参数
	var msg proto.Message
	if hand.Create != nil {
		dat := hand.Create()
		proto.Unmarshal(data, dat)
		msg = dat
	}
	_, err := hand.Handle(msg)
	if err != nil {
		xlog.Error(err.Error())
	}
}

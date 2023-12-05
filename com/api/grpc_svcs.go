package api

import (
	"errors"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"lark/pb"

	"google.golang.org/protobuf/proto"
)

type SvcClienter interface {
	SendMsg(pb.APIMsgId, proto.Message, proto.Message) error
	PostMsg(pb.APIMsgId, proto.Message) error
}

var svc map[pb.ServerType]SvcClienter

func NewGrpcService(typ pb.ServerType) SvcClienter {
	if svc == nil {
		svc = map[pb.ServerType]SvcClienter{}
	}
	if c, ok := svc[typ]; ok && c != nil {
		return c
	}
	cli := NewGrpcClient(App.SysCfg.Servers[uint32(typ)])
	svc[typ] = cli
	return cli
}

// 发送消息 等待返回
func SendMsg(typ pb.ServerType, id pb.APIMsgId, data proto.Message, res proto.Message) error {
	cli, ok := svc[typ]
	if !ok || cli == nil {
		return errors.New(utils.CONST_GRPC_SERVICE_NONE)
	}
	return cli.SendMsg(id, data, res)
}

// 发送消息 没有返回
func PostMsg(typ pb.ServerType, id pb.APIMsgId, data proto.Message) error {
	return SendMsg(typ, id, data, nil)
}

// 发送给所有注册的客户端
// 注意：如果该服务启动了多个实例，只有其中一个会收到消息
func PostToAll(id pb.APIMsgId, data proto.Message) {
	for _, cli := range svc {
		err := cli.PostMsg(id, data)
		if err != nil {
			xlog.Warn(err.Error())
		}
	}
}

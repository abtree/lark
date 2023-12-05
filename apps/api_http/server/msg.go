package server

import (
	"lark/com/pkgs/xsync"
	"lark/pb"

	"google.golang.org/protobuf/proto"
)

// 通用消息处理 所有server都可以用
func (s *apiServer) RegisterCommonMsg() {
	s.grpcServer.Register(0, uint32(pb.APIMsgId_BackToWeb),
		func() proto.Message {
			return &pb.WebProto{}
		},
		s.WebMsg,
	)
}

func (s *apiServer) WebMsg(msg proto.Message) (proto.Message, error) {
	task := msg.(*pb.WebProto)
	// ctrl.WebCtrl.Sign(task.Guid, task)
	xsync.Sign(task.Guid, task)
	return nil, nil
}

package server

import (
	"lark/pb"

	"google.golang.org/protobuf/proto"
)

// 通用消息处理 所有server都可以用
func (s *authServer) RegisterCommonMsg() {
	s.Register(0, uint32(pb.APIMsgId_EAuth),
		func() proto.Message {
			return &pb.AuthProto{}
		},
		s.Auth,
	)
}

func (s *authServer) Auth(msg proto.Message) (proto.Message, error) {
	return s.authService.OnProto(msg.(*pb.AuthProto))
}

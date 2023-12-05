package server

import (
	"lark/com/pkgs/xsync"
	"lark/pb"

	"google.golang.org/protobuf/proto"
)

// 通用消息处理 所有server都可以用
func (s *giftpackServer) RegisterCommonMsg() {
	s.Register(0, uint32(pb.APIMsgId_EGiftPack),
		func() proto.Message {
			return &pb.GiftPackProto{}
		},
		s.Giftpack,
	)
}

func (s *giftpackServer) Giftpack(msg proto.Message) (proto.Message, error) {
	guid := xsync.GetGuid()
	s.gpService.PostTask(guid, msg.(*pb.GiftPackProto))
	res, err := xsync.Wait(guid)
	return res.(*pb.GiftPackProto), err
}

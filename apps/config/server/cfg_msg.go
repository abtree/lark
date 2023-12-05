package server

import (
	"lark/pb"

	"google.golang.org/protobuf/proto"
)

// 通用消息处理 所有server都可以用
func (s *configServer) RegisterCommonMsg() {
	s.Register(0, uint32(pb.APIMsgId_GetAllCfg),
		nil,
		s.GetAllCfg,
	)
}

func (s *configServer) GetAllCfg(msg proto.Message) (proto.Message, error) {
	return s.cfgService.GetAllCfg(), nil
}

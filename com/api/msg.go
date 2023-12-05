package api

import "google.golang.org/protobuf/proto"

type IMsgServer interface {
	RegisterMsg()
}

type MsgServer struct {
	//处理grpc消息 [serverType msgID handler]
	handles map[uint32]map[uint32]*MsgHandle
}

// 创建消息参数接收对象
type SrvCreateFunc func() proto.Message

// 消息处理函数
type SrvHandleFunc func(msg proto.Message) (proto.Message, error)

// 消息处理函数注册
type MsgHandle struct {
	Create SrvCreateFunc
	Handle SrvHandleFunc
}

// 注册消息处理函数
func (s *MsgServer) Register(typ, id uint32, create SrvCreateFunc, hand SrvHandleFunc) {
	if s.handles == nil {
		s.handles = map[uint32]map[uint32]*MsgHandle{}
	}
	hs, ok := s.handles[typ]
	if !ok || hs == nil {
		hs = map[uint32]*MsgHandle{}
		s.handles[typ] = hs
	}
	hs[id] = &MsgHandle{
		Create: create,
		Handle: hand,
	}
}

//获取消息处理对象
func (s *MsgServer) getHandler(typ, id uint32) *MsgHandle {
	hs, ok := s.handles[typ]
	if !ok || hs == nil {
		return nil
	}
	var hand *MsgHandle
	hand, ok = hs[id]
	if !ok || hand == nil {
		return nil
	}
	return hand
}

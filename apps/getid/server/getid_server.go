package server

import (
	"lark/apps/getid/service"
	"lark/com/api"
	"lark/com/pkgs/xmongo"
	"lark/com/pkgs/xsync"
	"lark/pb"

	"google.golang.org/protobuf/proto"
)

type GetidServer interface {
	api.IGrpcServer
}

type getidServer struct {
	//集成grpc服务
	api.GrpcServer
	//逻辑处理service
	svc *service.GetidService
}

func NewGetidServer(svc *service.GetidService) GetidServer {
	return &getidServer{
		svc: svc,
	}
}

func (s *getidServer) Init() error {
	xmongo.NewMongoClient(api.App.SysCfg.Mongo)

	s.svc.Init()
	return nil
}

// 注册消息处理函数
func (s *getidServer) RegisterMsg() {
	s.Register(0, uint32(pb.APIMsgId_EGetid), nil, s.GetId)
}

func (s *getidServer) GetId(msg proto.Message) (proto.Message, error) {
	guid := xsync.GetGuid()
	s.svc.PostTask(guid)
	ret, err := xsync.Wait(guid)
	if err != nil {
		return nil, err
	}
	return ret.(*pb.GetidProto), nil
}

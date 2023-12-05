package api

import (
	"errors"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"lark/pb"

	"google.golang.org/protobuf/proto"
)

func grpcError(code pb.ServerError, id uint32, msg string) (*pb.SvcResp, error) {
	resp := &pb.SvcResp{}
	resp.Code = int32(code)
	resp.Msg = msg
	resp.ID = id
	return resp, errors.New(utils.CONST_GRPC_NOT_HANDLE)
}

// 处理消息
func (s *GrpcServer) Handle(typ uint32, req *pb.SvcReq) (*pb.SvcResp, error) {
	if s.handles == nil {
		xlog.Error(utils.CONST_GRPC_NOT_HANDLE)
		return grpcError(pb.ServerError_GrpcNotRegisterHandle, req.ID, utils.CONST_GRPC_NOT_HANDLE)
	}
	//获取专有消息处理函数
	hand := s.getHandler(typ, req.ID)
	if hand == nil {
		//获取通用消息处理函数
		hand = s.getHandler(0, req.ID)
	}
	if hand == nil {
		xlog.Error(utils.CONST_GRPC_NOT_HANDLE)
		return grpcError(pb.ServerError_GrpcNotRegisterHandle, req.ID, utils.CONST_GRPC_NOT_HANDLE)
	}

	//解析请求参数
	var msg proto.Message
	if hand.Create != nil {
		dat := hand.Create()
		proto.Unmarshal(req.Data, dat)
		msg = dat
	}
	//处理消息
	back, err := hand.Handle(msg)
	resp := &pb.SvcResp{}
	resp.ID = req.ID
	if err != nil {
		resp.Code = int32(pb.ServerError_GrpcHandleMsgError)
		resp.Msg = err.Error()
		xlog.Error(err.Error())
		return resp, err
	}
	//处理返回数据
	if back != nil {
		byts, err := proto.Marshal(back)
		if err != nil {
			xlog.Error(err.Error())
			resp.Code = int32(pb.ServerError_GrpcPackMsgError)
			resp.Msg = err.Error()
			return resp, err
		}
		resp.Data = byts
	}
	return resp, nil
}

package service

import (
	"errors"
	"lark/com/api"
	"lark/com/pkgs/xjwt"
	"lark/com/utils"
	"lark/pb"

	"google.golang.org/protobuf/proto"
)

// type AuthService interface {
// 	SayHello(*pb.UserInfo) (*pb.MsgStr, error)
// }

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) OnProto(req *pb.AuthProto) (resp *pb.AuthProto, err error) {
	switch req.Op {
	case pb.AuthProto_Register:
		resp, err = s.Register(req)
	case pb.AuthProto_Login:
		resp, err = s.Register(req)
	case pb.AuthProto_LoginToken:
		resp, err = s.Register(req)
	default:
		err = errors.New(utils.Const_PB_No_Operator)
	}
	return
}

// 新用户注册
func (s *AuthService) Register(req *pb.AuthProto) (*pb.AuthProto, error) {
	id := &pb.GetidProto{}
	err := api.SendMsg(pb.ServerType_GetId, pb.APIMsgId_EGetid, nil, id)
	if err != nil { //获取id失败 无法创建用户
		return nil, err
	}
	req.User.Id = uint64(id.Id)
	byts, err := proto.Marshal(req.User)
	if err != nil { //编码用户数据失败（一般不会发生）
		return nil, err
	}
	t, err := xjwt.CreateToken(byts, false, 86400)
	if err != nil { //创建Token失败
		return nil, err
	}
	req.Token = t.Token
	req.Expire = int32(t.Expire)
	return req, nil
}

// 登录
func (s *AuthService) Login(req *pb.AuthProto) (*pb.AuthProto, error) {
	return nil, nil
}

// 使用Token登录
func (s *AuthService) LoginToken(req *pb.AuthProto) (*pb.AuthProto, error) {
	return nil, nil
}

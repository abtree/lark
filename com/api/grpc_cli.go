package api

import (
	"context"
	"errors"
	"lark/com/pkgs/xgrpc"
	"lark/com/pkgs/xlog"
	"lark/pb"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type GrpcClient struct {
	opt *xgrpc.ClientDialOption
}

func NewGrpcClient(srv *pb.CfgServers) *GrpcClient {
	cfg := App.SysCfg
	opt := xgrpc.NewClientDialOption(cfg.Etcd, srv, cfg.Jaeger, App.SrvCfg.Name)
	return &GrpcClient{
		opt: opt,
	}
}

func (c *GrpcClient) SrvSvc(req *pb.SvcReq) *pb.SvcResp {
	conn := c.opt.GetClientConn()
	client := pb.NewSrvServiceClient(conn)
	mtdata := metadata.New(map[string]string{ServerType: App.Envs[ServerType]})
	ctx := metadata.NewOutgoingContext(context.Background(), mtdata)
	resp, err := client.SrvSvc(ctx, req)
	if err != nil {
		xlog.Warn(err.Error())
		return &pb.SvcResp{
			Msg: err.Error(),
		}
	}
	return resp
}

func (c *GrpcClient) SendMsg(id pb.APIMsgId, data proto.Message, res proto.Message) error {
	dat := []byte{}
	var err error
	if data != nil {
		dat, err = proto.Marshal(data)
		if err != nil {
			return err
		}
	}
	return c.SendBytes(id, dat, res)
}

func (c *GrpcClient) SendBytes(id pb.APIMsgId, data []byte, res proto.Message) error {
	msg := &pb.SvcReq{
		ID:   uint32(id),
		Data: data,
	}
	resp := c.SrvSvc(msg)
	if resp.Msg != "" {
		return errors.New(resp.Msg)
	}
	if res != nil {
		return proto.Unmarshal(resp.Data, res)
	} else {
		return nil
	}
}

func (c *GrpcClient) PostMsg(id pb.APIMsgId, data proto.Message) error {
	return c.SendMsg(id, data, nil)
}

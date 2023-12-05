package main

import (
	"context"
	"lark/com/pkgs/xgrpc"
	"lark/com/pkgs/xloadcfg"
	"lark/com/pkgs/xlog"
	"lark/pb"
)

func main() {
	cfg := &pb.MsgSysConfigs{}
	xloadcfg.Run("../../../configs", cfg, nil)
	xlog.Shared(cfg.Logger, "examples")

	svrcfg := cfg.Servers[2] //lark_config
	svrcfg.Host = "127.0.0.1"

	opt := xgrpc.NewClientDialOption(cfg.Etcd, svrcfg, cfg.Jaeger, "test")
	conn := opt.GetClientConn()
	client := pb.NewSrvServiceClient(conn)

	req := &pb.SvcReq{
		ID: 1,
	}
	resp, err := client.SrvSvc(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	} else {
		xlog.Debug(resp)
	}
}

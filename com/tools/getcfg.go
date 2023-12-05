package tools

import (
	"lark/com/api"
	"lark/com/pkgs/xlog"
	"lark/pb"
)

func GetAllCfg() (*pb.MsgAllConfigs, error) {
	cfg := &pb.MsgAllConfigs{}
	err := api.SendMsg(pb.ServerType_Config,
		pb.APIMsgId_GetAllCfg, nil, cfg)
	if err != nil {
		xlog.Warn(err.Error())
		return nil, err
	}
	return cfg, nil
}

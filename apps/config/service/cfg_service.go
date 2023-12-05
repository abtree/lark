package service

import (
	"lark/com/pkgs/xloadcfg"
	"lark/pb"
	"sync/atomic"
)

// type CfgService interface {
// 	LoadAllConfigs()
// 	GetAllCfg() *pb.MsgAllConfigs
// }

type CfgService struct {
	allcfg atomic.Value
}

func NewCfgService() *CfgService {
	return &CfgService{}
}

func (t *CfgService) LoadAllConfigs() {
	dp := &pb.MsgAllConfigs{
		Configs: &pb.MsgConfigs{},
		Yyacts:  &pb.MsgYYactConfigs{},
	}
	xloadcfg.Run("./files", dp.Configs, dp.Yyacts)
	t.allcfg.Store(dp)
	// xlog.Debug(t.allcfg)
}

func (t *CfgService) GetAllCfg() *pb.MsgAllConfigs {
	c := t.allcfg.Load()
	return c.(*pb.MsgAllConfigs)
}

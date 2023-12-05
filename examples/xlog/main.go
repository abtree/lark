package main

import (
	"lark/com/pkgs/xloadcfg"
	"lark/com/pkgs/xlog"
	"lark/pb"
)

func main() {
	cfg := &pb.MsgSysConfigs{}
	xloadcfg.Run("../../configs", cfg, nil)

	xlog.Shared(cfg.Logger, "examples")
	xlog.Debug("Debug")
	xlog.Info("Info")
	xlog.Warn("Warn")
	xlog.Error("Error")
	// xlog.Panic("Panic")
}

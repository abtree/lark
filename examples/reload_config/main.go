package main

import (
	"fmt"
	"lark/com/pkgs/xloadcfg"
	"lark/pb"
)

func main() {
	cfg := &pb.MsgSysConfigs{}

	xloadcfg.Run("../../configs", cfg, nil)
	fmt.Println(cfg)
}

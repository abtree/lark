package main

import (
	"fmt"
	"lark/com/pkgs/xloadcfg"
	"lark/com/pkgs/xlog"
	"lark/com/pkgs/xredis"
	"lark/pb"
	"time"
)

func main() {
	cfg := &pb.MsgSysConfigs{}
	xloadcfg.Run("../../configs", cfg, nil)
	xlog.Shared(cfg.Logger, "examples")

	xredis.NewRedisClient(cfg.Redis)
	xredis.Set("test", "test", 10*time.Second)
	fmt.Println(xredis.TTL("test"))
	s, e := xredis.Get("test")
	if e == nil {
		fmt.Println(s)
	}
	s, e = xredis.Get("test111")
	if e != nil {
		fmt.Println(e.Error())
	}
	xredis.Del("test")
}

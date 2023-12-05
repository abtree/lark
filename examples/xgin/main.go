package main

import (
	"fmt"
	"lark/com/pkgs/xgin"
	"lark/com/pkgs/xjwt"
	"lark/com/pkgs/xloadcfg"
	"lark/com/pkgs/xlog"
	"lark/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := &pb.MsgSysConfigs{}
	xloadcfg.Run("../../configs", cfg, nil)
	xlog.Shared(cfg.Logger, "examples")

	engine := xgin.NewServer()
	engine.Use(JwtAuth(), Test())
	engine.Engine.GET("/", Hello)
	engine.Run(8080)
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		str := xgin.ParseFromHeader(ctx)
		if str == "" {
			var err error
			str, err = xgin.ParseFromCookie(ctx)
			if err != nil {
				ctx.Abort()
				ctx.SecureJSON(http.StatusForbidden, "Token 验证失败")
				return
			}
		}
		t, err := xjwt.ParseFromToken(str)
		if err != nil {
			ctx.Abort()
			ctx.SecureJSON(http.StatusForbidden, "Token 验证失败")
			return
		}
		fmt.Println("Token 验证成功", t)
	}
}

func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("测试中间件")
	}
}

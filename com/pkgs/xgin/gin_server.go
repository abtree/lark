package xgin

import (
	"lark/com/pkgs/xlog"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	Engine *gin.Engine
}

// create a new gin server
func NewServer() *GinServer {
	svr := &GinServer{}
	gin.SetMode(gin.ReleaseMode)
	svr.Engine = gin.New()
	return svr
}

//添加中间件
func (s *GinServer) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.Use(middleware...)
}

//启动GinServer监听
func (s *GinServer) Run(port int) {
	addr := ":" + strconv.Itoa(port)
	xlog.Info("gin server start with", addr)
	err := s.Engine.Run(addr)
	if err != nil {
		xlog.Error("GinServer start error", err.Error())
	}
}

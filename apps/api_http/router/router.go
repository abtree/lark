package router

import (
	"lark/apps/api_http/ctrl"

	"github.com/gin-gonic/gin"
)

// 注册路由
func Register(engine *gin.Engine) {
	//注册静态资源目录
	engine.Static("/", "./web")
	//开放路由组
	registerOpenRouter(engine)
	//需要授权路由组
	registerApiRouter(engine)
}

// 注册开放路由
func registerOpenRouter(engine *gin.Engine) {
	openGroup := engine.Group("open")
	registerAuthRouter(openGroup)
}

func registerAuthRouter(group *gin.RouterGroup) {
	authGroup := group.Group("auth")
	authGroup.POST("register", ctrl.WebCtrl.Auth)
	authGroup.POST("login", ctrl.WebCtrl.Auth)
	authGroup.POST("login_token", ctrl.WebCtrl.Auth)
}

// 注册授权路由
func registerApiRouter(engine *gin.Engine) {
	apiGroup := engine.Group("api")
	//由于有权限认证，该认证用不上
	// apiGroup.Use(JwtAuth())
	// apiGroup.Use(OperatorLog())
	registerGiftpackRouter(apiGroup)
}

func registerGiftpackRouter(group *gin.RouterGroup) {
	gpGroup := group.Group("giftpack", ctrl.WebCtrl.Privilage(1, false))
	gpGroup.POST("create", ctrl.WebCtrl.GiftPack)
	gpGroup.POST("update", ctrl.WebCtrl.GiftPack)
	gpGroup.POST("download", ctrl.WebCtrl.GiftPack)
}

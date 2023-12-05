package router

import "github.com/gin-gonic/gin"

//token验证的中间件
func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

//日志记录(需要记录操作日志时使用)
func OperatorLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

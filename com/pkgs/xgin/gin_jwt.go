package xgin

import "github.com/gin-gonic/gin"

//从cookie读取jwt
func ParseFromCookie(ctx *gin.Context) (string, error) {
	return ctx.Cookie("jwt")
}

//从header读取jwt
func ParseFromHeader(ctx *gin.Context) string {
	jwt := ctx.GetHeader("token")
	if jwt == "" {
		jwt = ctx.Query("token")
	}
	return jwt
}

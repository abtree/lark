package xgin

import (
	"lark/com/pkgs/xlog"
	"lark/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Msg  string `json:"msg"`
	Code int32  `json:"code"`
	Data interface{}
}

func ParamJson(ctx *gin.Context, pars interface{}) bool {
	if err := ctx.BindJSON(pars); err != nil {
		xlog.Warn(err.Error())
		Error(ctx, int32(pb.ServerError_ParamError), err.Error())
		return false
	}
	return true
}

func (r *Resp) SetResult(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func Success(ctx *gin.Context, data interface{}) {
	if data == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func Error(ctx *gin.Context, code int32, err string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err,
	})
}

func Result(ctx *gin.Context, resp *Resp) {
	if resp.Code > 0 {
		Error(ctx, resp.Code, resp.Msg)
	} else {
		Success(ctx, resp.Data)
	}
}

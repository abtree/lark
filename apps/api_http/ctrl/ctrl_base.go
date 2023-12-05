package ctrl

import (
	"lark/com/api"
	"lark/com/pkgs/xgin"
	"lark/com/pkgs/xlog"
	"lark/com/pkgs/xsync"
	"lark/com/utils"
	"lark/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

type Controller struct {
}

var WebCtrl = &Controller{}

func (ctrl *Controller) GetPars(ctx *gin.Context, pars proto.Message) bool {
	if pars != nil {
		ok := xgin.ParamJson(ctx, pars)
		if !ok {
			xlog.Warn(utils.CONST_GIN_PARSE_JSON)
			xgin.Error(ctx, int32(pb.ServerError_GinServerError), utils.CONST_GIN_PARSE_JSON)
			return false
		}
	}
	return true
}

// 同步的请求
func (ctrl *Controller) Svc(ctx *gin.Context, typ pb.ServerType, id pb.APIMsgId, pars proto.Message, res proto.Message) {
	ok := ctrl.GetPars(ctx, pars)
	if !ok {
		return
	}
	err := api.SendMsg(typ, id, pars, res)
	if err != nil {
		xlog.Warn(err.Error())
		xgin.Error(ctx, int32(pb.ServerError_GinServerError), err.Error())
		return
	}
	if res != nil {
		w := &xgin.Resp{
			Data: res,
		}
		xgin.Result(ctx, w)
	}
}

/*
异步的请求 需要等待返回消息(不推荐使用)
调用 router := xsync.GetGuid() 获取router，并且router需要随消息传递
当有多个实例时，需要保证返回消息调用到正确的实例
*/
func (ctrl *Controller) ASync(ctx *gin.Context, typ pb.ServerType, id pb.APIMsgId, pars proto.Message, router int64, res proto.Message) {
	ok := ctrl.GetPars(ctx, pars)
	if !ok {
		return
	}
	err := api.PostMsg(typ, id, pars)
	if err != nil {
		xlog.Warn(err.Error())
		xgin.Error(ctx, int32(pb.ServerError_GinServerError), err.Error())
	}
	if res != nil {
		task, err := xsync.Wait(router)
		w := &xgin.Resp{}
		if err != nil {
			w.Code = int32(pb.ServerError_GinServerError)
			w.Msg = err.Error()
		} else {
			err = proto.Unmarshal(task.(*pb.WebProto).Data, res)
			if err != nil {
				w.Code = int32(pb.ServerError_GinServerError)
				w.Msg = err.Error()
			} else {
				w.Data = res
			}
		}
		xgin.Result(ctx, w)
	}
}

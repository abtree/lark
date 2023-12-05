package ctrl

import (
	"lark/com/pkgs/xgin"
	"lark/com/pkgs/xjwt"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"lark/pb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

/*
权限认证(同一个group下，权限相同)
如果有特殊的需求，需要写在相应的请求下
*/
func (ctrl *Controller) Privilage(priv uint32, force bool) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		str := xgin.ParseFromHeader(ctx)
		if str == "" {
			ctx.Redirect(http.StatusFound, "/open/hello")
			return
		}
		t, err := xjwt.Decode(str)
		if err != nil {
			xlog.Warn(err.Error())
			ctx.Redirect(http.StatusFound, "/open/hello")
			return
		}
		if time.Now().Unix() > t.Expire {
			ctx.Redirect(http.StatusFound, "/open/hello")
			return
		}
		user := &pb.UserInfo{}
		proto.Unmarshal(t.UserData, user)
		b := false
		for _, v := range user.Privilage {
			if v == priv {
				b = true
				break
			}
		}
		if !b {
			ctx.Abort()
			ctx.SecureJSON(http.StatusForbidden, utils.Const_Token_Forbidden)
			return
		}
		if force {
			// todo: 向AuthServer请求验证
		}
	}
}

func (ctrl *Controller) Auth(ctx *gin.Context) {
	pars := &pb.AuthProto{}
	back := &pb.AuthProto{}
	ctrl.Svc(ctx, pb.ServerType_Auth, pb.APIMsgId_EAuth, pars, back)
}

package ctrl

import (
	"lark/pb"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GiftPack(ctx *gin.Context) {
	pars := &pb.GiftPackProto{}
	back := &pb.GiftPackProto{}
	ctrl.Svc(ctx, pb.ServerType_Giftpack, pb.APIMsgId_EGiftPack, pars, back)
}

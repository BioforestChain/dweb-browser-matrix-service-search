package userController

import (
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/entity/req/userReqEntity"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error/common"
	baseController "github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/http/controller"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/logic/userLogic"
	"github.com/gin-gonic/gin"
)

type controller struct {
	baseController.BaseController
}

func NewController(ctx *gin.Context) *controller {
	return &controller{baseController.NewBaseController(ctx)}
}
func (c *controller) UserList() {
	req := userReqEntity.List{}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Fail(common.ReqParamErr, c.GetValidMsg(err, &req))
		return
	}

	logic := userLogic.NewLogic(c.GCtx)
	resp, err := logic.GetUserList(req)
	if err != nil {
		c.Fail(err.Code(), err.Message())
		return
	}
	c.Success(resp)
}

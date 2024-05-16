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

//1. 查缓存,获取钱包地址对用的profile信息
//2. 并发调用synapse接口查询用户信息  (https://172.25.11.243/_matrix/client/v3/user_directory/search)
//3.

//{"search_term":"abc","limit":50}
//	{
//	   "limited": false,
//	   "results": [
//	       {
//	           "user_id": "@bagen008:172.25.11.243",
//	           "display_name": "bagen008",
//	           "avatar_url": null,
//	           "wallet_address": "abc"
//	       }
//	   ]
//	}

func (c *controller) UserInfo() {
	req := userReqEntity.Info{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Fail(common.ReqParamErr, c.GetValidMsg(err, &req))
		return
	}

	logic := userLogic.NewLogic(c.GCtx)
	resp, err := logic.GetUserInfo(req)
	if err != nil {
		c.Fail(err.Code(), err.Message())
		return
	}
	c.Success(resp)
}

//func (c *controller) UserList() {
//	req := userReqEntity.List{}
//	if err := c.ShouldBindQuery(&req); err != nil {
//		c.Fail(common.ReqParamErr, c.GetValidMsg(err, &req))
//		return
//	}
//
//	logic := userLogic.NewLogic(c.GCtx)
//	resp, err := logic.GetUserList(req)
//	if err != nil {
//		c.Fail(err.Code(), err.Message())
//		return
//	}
//	c.Success(resp)
//}

package userRouter

import (
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/http/controller/api/userController"
	"github.com/gin-gonic/gin"
)

func RegisterUser(router *gin.Engine) {

	router.POST("/user-info", func(ctx *gin.Context) {
		userController.NewController(ctx).UserInfo()
	})
	//router.GET("/user-list", func(ctx *gin.Context) {
	//	userController.NewController(ctx).UserList()
	//})
}

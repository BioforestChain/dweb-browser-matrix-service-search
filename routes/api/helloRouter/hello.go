package helloRouter

import (
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/http/controller/api/helloController"
	"github.com/gin-gonic/gin"
)

func RegisterHello(router *gin.Engine) {

	router.GET("/test", func(ctx *gin.Context) {
		helloController.NewController(ctx).Hello()
	})
}

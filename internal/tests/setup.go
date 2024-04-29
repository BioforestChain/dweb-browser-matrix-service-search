package tests

import (
	"context"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/model/common"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/pkg/support-go/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"testing"
)

var DB *gorm.DB

func Setup() {
	testing.Init()
	bootstrap.DevEnv = bootstrap.EnvLocal
	bootstrap.Init()
	bootstrap.InitWeb([]gin.HandlerFunc{})
	DB = common.ConnectionObject(context.TODO()).DB
}

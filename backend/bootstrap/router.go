package bootstrap

import (
	"fast-image/global"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

var Route func(e *gin.Engine)

func InitRouter() {
	if global.GetAppMode() == global.APP_MODE_PROD {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	engine = gin.Default()

	if Route != nil {
		Route(engine)
	}
	global.Logger.Info("初始化ROUTER完成")
}

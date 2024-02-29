package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {

	engine := gin.Default()

	apiGroup := engine.Group("api")
	apiRouterGroup := RouterGroup{apiGroup}
	apiRouterGroup.SettingsRouter()

	return engine

}

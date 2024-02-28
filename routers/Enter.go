package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//mode := global.Config.System.Mode
	//gin.SetMode(mode)
	engine := gin.Default()

	engine.GET("hello", func(context *gin.Context) {
		context.JSON(200, "hello")

	})

	return engine

}

package router

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "retail/docs"

	"github.com/gin-gonic/gin"
	"retail/handler"
)

func InitRouter() *gin.Engine {

	engine := gin.New()
	engine.Use(gin.Logger())

	// r.GET("/login", api.GetAuth)

	engine.Use(gin.Recovery())

	// r.GET("/services/:id", controller.OneGoods)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api/v1")
	api.Use()
	{

		engine.GET("/index", handler.Index)

	}

	return engine
}

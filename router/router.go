package router

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "retail/docs"

	"github.com/gin-gonic/gin"
	"retail/handler/coupon"
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

		// engine.GET("/conpons", handler.Index)
		// engine.GET("/conpons/:id", handler.Index)
		// engine.POST("/conpons", handler.Index)
		// engine.PUT("/conpons", handler.Index)
		// engine.DELETE("/conpons/:id",handler.Index)


		engine.GET("/conpons/type", coupon.GetType)
		engine.GET("/conpons/type/:id", coupon.GetTypeOne)
		engine.POST("/conpons/type", coupon.CreateType)
		engine.PUT("/conpons/type/:id", coupon.ModifyType)
		engine.DELETE("/conpons/type/:id",coupon.DelType)

	}

	return engine
}

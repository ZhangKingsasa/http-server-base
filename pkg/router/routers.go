package router

import (
	"git.garena.com/shopee/MLP/aip/platform/aip-user-service/pkg/controller"
	"git.garena.com/shopee/MLP/aip/platform/aip-user-service/pkg/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetRouter() *gin.Engine {
	router := gin.New()
	// common middleware
	router.Use(middleware.Logger(), middleware.Cors())

	// api doc
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// health check
	router.GET("/healthz", controller.Healthz)

	return router
}
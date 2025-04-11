package router

import (
	"github.com/gin-gonic/gin"

	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/handler"
	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	api := router.Group("/api/v1")
	{
		handler.RegisterHealthCheckerRoutes(api)
		handler.RegisterOrderRoutes(api)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

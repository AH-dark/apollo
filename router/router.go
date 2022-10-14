package router

import (
	"github.com/AH-dark/apollo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.RequestId())

	api := router.Group("/api/")
	BuildApi(api)

	return router
}

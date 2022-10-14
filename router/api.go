package router

import (
	"github.com/AH-dark/apollo/controller"
	"github.com/AH-dark/apollo/controller/auth"
	"github.com/AH-dark/apollo/middleware"
	"github.com/gin-gonic/gin"
)

func BuildApi(r *gin.RouterGroup) {
	r.Use(middleware.CORS())
	r.Use(middleware.Session())
	r.Use(middleware.Auth())

	r.GET("ping", controller.PingHandler)

	auths := r.Group("auth")
	{
		auths.POST("login", auth.LoginHandler)
	}
}

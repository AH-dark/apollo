package router

import (
	"github.com/AH-dark/apollo/controller"
	"github.com/AH-dark/apollo/controller/auth"
	"github.com/AH-dark/apollo/controller/info"
	"github.com/AH-dark/apollo/controller/session"
	"github.com/AH-dark/apollo/middleware"
	"github.com/gin-gonic/gin"
)

func BuildApi(r *gin.RouterGroup) {
	r.Use(middleware.CORS())
	r.Use(middleware.Session())
	r.Use(middleware.Auth())

	r.GET("ping", controller.PingHandler)

	infos := r.Group("info")
	{
		// /api/info/site GET 获取站点信息
		infos.GET("site", info.SiteInfoBasicHandler)
	}

	auths := r.Group("auth")
	{
		// /api/auth/login POST 登录
		auths.POST("login", auth.LoginHandler)

		// /api/auth/logout POST 登出
		auths.POST("logout", middleware.LoginOnly, auth.LogoutHandler)
	}

	sessions := r.Group("session")
	{
		// /api/session/me GET 获取当前用户信息
		sessions.GET("me", session.CurrentUserHandler)
	}
}

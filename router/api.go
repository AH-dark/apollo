package router

import (
	"github.com/AH-dark/apollo/controller"
	"github.com/AH-dark/apollo/controller/auth"
	"github.com/AH-dark/apollo/controller/comment"
	"github.com/AH-dark/apollo/controller/info"
	"github.com/AH-dark/apollo/controller/session"
	"github.com/AH-dark/apollo/middleware"
	"github.com/AH-dark/apollo/pkg/hashids"
	"github.com/gin-gonic/gin"
)

func BuildApi(r *gin.RouterGroup) {
	r.Use(middleware.CORS())
	r.Use(middleware.Session())
	r.Use(middleware.Auth())
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(200)
	})

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

	comments := r.Group("comment")
	{
		// /api/comment/submit POST 提交评论
		comments.POST("submit", comment.SubmitCommentHandler)

		// /api/comment/list GET 获取评论列表
		comments.GET("list", comment.ListCommentsHandler)

		// /api/comment/:cid GET 获取评论详情
		comments.GET(":cid", middleware.HashIdParser(hashids.CommentHash, "cid"), comment.GetCommentHandler)
	}
}

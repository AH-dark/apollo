package middleware

import (
	"time"

	"github.com/AH-dark/apollo/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func CORS() gin.HandlerFunc {
	c := cors.Config{
		AllowOrigins:     config.CORS.AllowOrigins,
		AllowMethods:     config.CORS.AllowMethods,
		AllowHeaders:     config.CORS.AllowHeaders,
		ExposeHeaders:    config.CORS.ExposeHeaders,
		AllowCredentials: config.CORS.AllowCredentials,
		MaxAge:           time.Duration(config.CORS.MaxAge) * time.Second,
	}

	if _, exist := lo.Find(c.AllowOrigins, func(origin string) bool {
		return origin == "*"
	}); exist {
		c.AllowAllOrigins = true
		c.AllowOrigins = nil
	}

	return cors.New(c)
}

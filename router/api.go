package router

import (
	"github.com/AH-dark/apollo/middleware"
	"github.com/gin-gonic/gin"
)

func BuildApi(r *gin.RouterGroup) {
	r.Use(middleware.CORS())
}

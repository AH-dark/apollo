package controller

import (
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingHandler ping
func PingHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("test", "test")
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.NewHttpError(
			http.StatusInternalServerError,
			err.Error(),
			c.MustGet("requestId").(string),
		))
		return
	}

	c.JSON(http.StatusOK, serializer.NewSuccessResponse("pong"))
}

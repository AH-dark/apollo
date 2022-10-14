package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return requestid.New(requestid.WithCustomHeaderStrKey("X-Apollo-Request-Id"), requestid.WithHandler(func(c *gin.Context, id string) {
		c.Set("request_id", id)
	}))
}

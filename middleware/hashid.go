package middleware

import (
	"github.com/AH-dark/apollo/pkg/hashids"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/gin-gonic/gin"
)

func HashIdParser(t hashids.HashIDType, paramKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cid := c.Param(paramKey)
		id, err := hashids.Decode(cid, t)
		if err != nil {
			c.AbortWithStatusJSON(400, serializer.NewHttpError(
				400,
				"invalid cid",
				c.MustGet("request_id").(string),
			))
			return
		}

		c.Set(paramKey, id)
		c.Next()
	}
}

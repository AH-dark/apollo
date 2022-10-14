package session

import (
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/AH-dark/apollo/vo"
	"github.com/gin-gonic/gin"
)

func CurrentUserHandler(c *gin.Context) {
	if s, ok := c.Get("user"); ok {
		if u := s.(*model.User); u != nil {
			c.JSON(200, serializer.NewSuccessResponse(vo.BuildUserVO(u)))
			return
		}
	}

	c.JSON(200, serializer.NewSuccessResponse(nil))
}

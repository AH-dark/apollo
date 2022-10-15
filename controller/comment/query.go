package comment

import (
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/AH-dark/apollo/vo"
	"github.com/gin-gonic/gin"
)

func GetCommentHandler(c *gin.Context) {
	cid := c.MustGet("cid").(uint)

	comment, err := model.Global.Comment.GetCommentByID(cid)
	if err != nil {
		c.AbortWithStatusJSON(400, serializer.NewHttpError(
			400,
			"invalid cid",
			c.MustGet("request_id").(string),
		))
		return
	}

	c.JSON(200, serializer.NewSuccessResponse(vo.BuildCommentVO(comment)))
}

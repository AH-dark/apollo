package comment

import (
	"github.com/AH-dark/apollo/dto"
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/AH-dark/apollo/vo"
	"github.com/gin-gonic/gin"
)

func SubmitCommentHandler(c *gin.Context) {
	var payload dto.SubmitCommentDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Log().
			WithField(log.FieldGinContext, c).
			WithError(err).
			Error("Bind json error")
		c.JSON(400, serializer.NewHttpError(400, err.Error(), c.MustGet("request_id").(string)))
		return
	}

	comment, err := model.Global.Comment.CreateCommentByDTO(dto.CommentDTO{
		SubmitCommentDTO: payload,
		IP:               c.ClientIP(),
		RequestID:        c.MustGet("request_id").(string),
	})
	if err != nil {
		log.Log().
			WithField(log.FieldGinContext, c).
			WithError(err).
			Error("Create comment error")
		c.JSON(500, serializer.NewHttpError(500, err.Error(), c.MustGet("request_id").(string)))
		return
	}

	c.JSON(200, serializer.NewSuccessResponse(vo.BuildCommentVO(comment)))
}

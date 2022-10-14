package comment

import (
	"github.com/AH-dark/apollo/dto"
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/AH-dark/apollo/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func ListCommentsHandler(c *gin.Context) {
	var payload dto.ListCommentsDTO
	if err := c.ShouldBindQuery(&payload); err != nil {
		log.Log().
			WithField(log.FieldGinContext, c).
			WithError(err).
			Error("Bind query error")
		c.JSON(400, serializer.NewHttpError(400, err.Error(), c.MustGet("request_id").(string)))
		return
	}

	before := time.Now()
	if payload.Before != nil {
		before = time.UnixMicro(*payload.Before)
	}

	comments, err := model.Global.Comment.ListCommentsByStatus(model.CommentStatusReplied, before, payload.PageSize)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Log().
			WithField(log.FieldGinContext, c).
			WithError(err).
			Error("List comments error")
		c.JSON(500, serializer.NewHttpError(500, err.Error(), c.MustGet("request_id").(string)))
		return
	}

	counts := model.Global.Comment.CountCommentsByStatus(model.CommentStatusReplied)

	c.JSON(200, serializer.NewSuccessResponse(vo.BuildCommentVOList(comments, counts, before)))
}

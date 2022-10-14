package auth

import (
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user_id")
	err := session.Save()
	if err != nil {
		log.Log().
			WithField(log.FieldGinContext, c).
			WithError(err).
			Error("Save session error")
		c.JSON(500, serializer.NewHttpError(500, err.Error(), c.MustGet("request_id").(string)))
		return
	}

	log.Log().
		WithField(log.FieldGinContext, c).
		WithField(log.FieldUsername, c.MustGet("user").(*model.User).Username).
		Info("User logout")

	c.JSON(200, serializer.NewSuccessResponse(nil))
}

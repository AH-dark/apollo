package auth

import (
	"github.com/AH-dark/apollo/dto"
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/AH-dark/apollo/vo"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var payload dto.LoginDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Log().
			WithField(log.FieldGinContext, c).
			WithError(err).
			Error("Bind json error")
		c.JSON(http.StatusBadRequest, serializer.NewHttpError(
			http.StatusBadRequest,
			err.Error(),
			c.MustGet("request_id").(string),
		))
		return
	}

	ok, user := model.Global.User.ComparePassword(payload.Login, payload.Password)
	if !ok {
		c.JSON(http.StatusUnauthorized, serializer.NewHttpError(
			http.StatusUnauthorized,
			"Invalid username or password",
			c.MustGet("request_id").(string),
		))
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	err := session.Save()
	if err != nil {
		log.Log().
			WithField(log.FieldGinContext, c).
			WithError(err).
			Error("Save session error")
		c.JSON(http.StatusInternalServerError, serializer.NewHttpError(
			http.StatusInternalServerError,
			err.Error(),
			c.MustGet("request_id").(string),
		))
		return
	}

	c.JSON(http.StatusOK, serializer.NewSuccessResponse(vo.BuildUserVO(user)))
}

package middleware

import (
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sess := session.Get("user_id")
		if sess == nil {
			return
		}

		if v, ok := sess.(uint); ok {
			user, err := model.Global.User.GetUserByID(v)
			if err != nil {
				log.Log().
					WithField(log.FieldGinContext, c).
					WithError(err).
					Error("Get user by id error")
				return
			}

			c.Set("user", user)
		}
	}
}

func LoginOnly(c *gin.Context) {
	if c.MustGet("user") == nil {
		c.AbortWithStatusJSON(403, serializer.NewHttpError(http.StatusForbidden, "Login required", c.MustGet("requestId").(string)))
		return
	}
}

func NoLoginOnly(c *gin.Context) {
	if c.MustGet("user") != nil {
		c.AbortWithStatusJSON(403, serializer.NewHttpError(http.StatusForbidden, "No login required", c.MustGet("requestId").(string)))
		return
	}
}

func AdminOnly(c *gin.Context) {
	if c.MustGet("user") == nil {
		c.AbortWithStatusJSON(403, serializer.NewHttpError(http.StatusForbidden, "Login required", c.MustGet("requestId").(string)))
		return
	}

	user := c.MustGet("user").(*model.User)
	if !user.IsAdmin() {
		c.AbortWithStatusJSON(403, serializer.NewHttpError(http.StatusForbidden, "Admin only", c.MustGet("requestId").(string)))
		return
	}
}

package middleware

import (
	"net/http"

	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	if _, ok := c.Get("user"); !ok {
		c.AbortWithStatusJSON(403, serializer.NewHttpError(http.StatusForbidden, "Login required", c.MustGet("request_id").(string)))
		return
	}
}

func NoLoginOnly(c *gin.Context) {
	if _, ok := c.Get("user"); ok {
		c.AbortWithStatusJSON(403, serializer.NewHttpError(http.StatusForbidden, "No login required", c.MustGet("request_id").(string)))
		return
	}
}

func AdminOnly(c *gin.Context) {
	sess, ok := c.Get("user")
	if !ok {
		c.AbortWithStatusJSON(403, serializer.NewHttpError(http.StatusForbidden, "Login required", c.MustGet("request_id").(string)))
		return
	}

	user, ok := sess.(*model.User)
	if !ok {
		c.AbortWithStatusJSON(500, serializer.NewHttpError(http.StatusInternalServerError, "Internal error", c.MustGet("request_id").(string)))
		return
	}

	if !user.IsAdmin() {
		c.AbortWithStatusJSON(403, serializer.NewHttpError(http.StatusForbidden, "Admin only", c.MustGet("request_id").(string)))
		return
	}
}

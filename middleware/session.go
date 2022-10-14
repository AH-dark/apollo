package middleware

import (
	"fmt"
	"github.com/samber/lo"
	"net/http"
	"strings"

	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store sessions.Store

func InitSession() {
	if config.Redis.Host != "" {
		var err error
		store, err = redis.NewStore(10, config.Redis.Network, fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port), config.Redis.Password, []byte(config.System.SessionSecret))
		if err != nil {
			log.Log().WithError(err).Error("Init redis session error")
			store = memstore.NewStore([]byte(config.System.SessionSecret))
		} else {
			return
		}
	} else {
		store = memstore.NewStore([]byte(config.System.SessionSecret))
	}

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 3,
		HttpOnly: true,
		Secure:   config.CORS.Https,
		SameSite: lo.Switch[string, http.SameSite](strings.ToLower(config.CORS.SameSite)).
			Case("lax", http.SameSiteLaxMode).
			Case("strict", http.SameSiteStrictMode).
			Case("none", http.SameSiteNoneMode).
			Default(http.SameSiteDefaultMode),
	})
}

func Session() gin.HandlerFunc {

	return sessions.Sessions("apollo-session", store)
}

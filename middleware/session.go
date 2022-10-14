package middleware

import (
	"fmt"
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
		} else {
			return
		}
	}

	store = cookie.NewStore([]byte(config.System.SessionSecret))
}

func Session() gin.HandlerFunc {
	return sessions.Sessions("apollo-session", store)
}

package bootstrap

import (
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApplicationArgs struct {
	ForceMigrate bool
}

func InitApplication(args ApplicationArgs) {
	log.Log()

	if config.System.Debug {
		log.GlobalLogger.SetLevel(logrus.DebugLevel)
		gin.SetMode(gin.DebugMode)
	} else {
		log.GlobalLogger.SetLevel(logrus.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}

	log.Log().Info("Init application")
	Init(args)
}

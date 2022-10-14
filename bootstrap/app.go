package bootstrap

import (
	"apollo/config"
	"apollo/pkg/log"
	"github.com/sirupsen/logrus"
)

func InitApplication() {
	if config.System.Debug {
		log.GlobalLogger.SetLevel(logrus.DebugLevel)
	}

	log.Log().Info("init application")
	Init()
}

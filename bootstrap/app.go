package bootstrap

import (
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/sirupsen/logrus"
)

func InitApplication() {
	if config.System.Debug {
		log.GlobalLogger.SetLevel(logrus.DebugLevel)
	}

	log.Log().Info("init application")
	Init()
}

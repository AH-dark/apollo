package bootstrap

import (
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/sirupsen/logrus"
)

type ApplicationArgs struct {
	ForceMigrate bool
}

func InitApplication(args ApplicationArgs) {
	if config.System.Debug {
		log.GlobalLogger.SetLevel(logrus.DebugLevel)
	}

	log.Log().Info("Init application")
	Init(args)
}

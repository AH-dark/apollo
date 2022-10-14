package bootstrap

import (
	"github.com/AH-dark/apollo/middleware"
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/cache"
	"github.com/AH-dark/apollo/pkg/log"
)

func Init(args ApplicationArgs) {
	log.Log().Info("Initialing model...")
	model.Init(args.ForceMigrate)

	log.Log().Info("Initialing cache...")
	cache.Init()

	log.Log().Info("Initialing session...")
	middleware.InitSession()
}

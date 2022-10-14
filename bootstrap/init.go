package bootstrap

import (
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/log"
)

func Init(args ApplicationArgs) {
	log.Log().Info("Initialing model...")
	model.Init(args.ForceMigrate)
}

package main

import (
	"flag"
	"github.com/AH-dark/apollo/bootstrap"
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/router"
)

var args = bootstrap.ApplicationArgs{
	ForceMigrate: false,
}

func init() {
	flag.BoolVar(&args.ForceMigrate, "force-migrate", false, "force migrate database")
	flag.Parse()

	bootstrap.InitApplication(args)
}

func main() {
	r := router.InitRouter()

	err := r.Run(config.System.Listen)
	if err != nil {
		log.Log().WithError(err).Error("run server error")
	}
}

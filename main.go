package main

import (
	"github.com/AH-dark/apollo/bootstrap"
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/router"
)

func init() {
	bootstrap.InitApplication()
}

func main() {
	r := router.InitRouter()

	err := r.Run(config.System.Listen)
	if err != nil {
		log.Log().WithError(err).Error("run server error")
	}
}

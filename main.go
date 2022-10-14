package main

import (
	"apollo/bootstrap"
	"apollo/config"
	"apollo/pkg/log"
	"apollo/router"
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

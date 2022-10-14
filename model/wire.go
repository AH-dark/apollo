//go:build wireinject
// +build wireinject

package model

import (
	"github.com/google/wire"
)

func BuildClient(forceMigrate bool) *Client {
	wire.Build(NewClient, initDB, NewSettingService, NewCommentService, NewUserService)
	return nil
}

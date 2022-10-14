// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package model

// Injectors from wire.go:

func BuildClient(forceMigrate bool) *Client {
	db := initDB(forceMigrate)
	modelSettingService := NewSettingService(db)
	modelCommentService := NewCommentService(db)
	client := NewClient(db, modelSettingService, modelCommentService)
	return client
}

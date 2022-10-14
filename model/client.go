package model

import "gorm.io/gorm"

type Client struct {
	db *gorm.DB

	Setting SettingService
	Comment CommentService
	User    UserService
}

func NewClient(db *gorm.DB, setting SettingService, comment CommentService, user UserService) *Client {
	return &Client{
		db:      db,
		Setting: setting,
		Comment: comment,
		User:    user,
	}
}

var Global *Client

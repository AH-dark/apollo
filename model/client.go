package model

import "gorm.io/gorm"

type Client struct {
	db *gorm.DB

	Setting SettingService
	Comment CommentService
}

func NewClient(db *gorm.DB, setting SettingService, comment CommentService) *Client {
	return &Client{
		db:      db,
		Setting: setting,
		Comment: comment,
	}
}

var Global *Client

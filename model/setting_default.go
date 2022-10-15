package model

import "github.com/AH-dark/apollo/config"

var defaultSettings = []Setting{
	{
		Name:  "version",
		Type:  SettingTypeSystem,
		Value: config.AppVersion,
	},
	{
		Name:  "site_name",
		Type:  SettingTypeBasic,
		Value: "Apollo",
	},
	{
		Name:  "site_url",
		Type:  SettingTypeBasic,
		Value: "http://localhost:8080",
	},
	{
		Name:  "site_description",
		Type:  SettingTypeBasic,
		Value: "A simple comment board",
	},
	{
		Name:  "comment_title",
		Type:  SettingTypeBasic,
		Value: "Leave your comments!",
	},
	{
		Name:  "comment_background_image",
		Type:  SettingTypeBasic,
		Value: "https://source.unsplash.com/random/1920x1080",
	},
}

package model

type Setting struct {
	Name  string      `gorm:"primary_key;unique_index;not null;size:255"`
	Type  SettingType `gorm:"not null;size:255;index"`
	Value string      `gorm:"not null;size:255"`
}

type SettingType string

const (
	SettingTypeSystem       SettingType = "system"
	SettingTypeBasic        SettingType = "basic"
	SettingTypeNotification SettingType = "notification"
)

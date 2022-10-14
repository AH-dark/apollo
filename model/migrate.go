package model

import (
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func needMigration(db *gorm.DB) bool {
	err := db.Model(&Setting{}).Where("name = ? AND value = ?", "version", config.AppVersion).First(&Setting{}).Error
	return err != nil
}

func migrate(db *gorm.DB, force bool) {
	if !force && !needMigration(db) {
		return
	}

	err := db.AutoMigrate(&Setting{}, &Comment{})
	if err != nil {
		log.Log().WithError(err).Error("failed to migrate setting")
		return
	}

	addDefaultSettings(db)
}

func addDefaultSettings(db *gorm.DB) {
	var existSettings []Setting
	err := db.Find(&existSettings).Error
	if err != nil {
		log.Log().WithError(err).Error("failed to get settings")
		return
	}

	var settings = defaultSettings
	settings = lo.Filter(settings, func(setting Setting, i int) bool {
		_, exist := lo.Find(existSettings, func(existSetting Setting) bool {
			return existSetting.Name == setting.Name
		})

		return !exist
	})

	for _, setting := range settings {
		err := db.Create(&setting).Error
		if err != nil {
			log.Log().WithError(err).Error("failed to add default setting")
		}
	}
}

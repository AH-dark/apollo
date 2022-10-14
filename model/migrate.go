package model

import (
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/crypto"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func needMigration(db *gorm.DB) bool {
	err := db.Model(&Setting{}).Where("name = ? AND value = ?", "version", config.AppVersion).First(&Setting{}).Error
	return err != nil
}

func migrate(db *gorm.DB, force bool) {
	if !force && !needMigration(db) {
		log.Log().Info("no need to migrate")
		return
	}

	log.Log().Info("start migrating")
	defer log.Log().Info("end migrating")

	err := db.AutoMigrate(&Setting{}, &Comment{}, &User{})
	if err != nil {
		log.Log().WithError(err).Error("failed to migrate setting")
		return
	}

	addDefaultSettings(db)
	addDefaultUser(db)
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

func addDefaultUser(db *gorm.DB) {
	var user User
	err := db.Model(&User{}).Where("username = ?", "admin").First(&user).Error
	if err == nil {
		return
	}

	pass := util.RandString(8)

	err = db.Create(&User{
		Username: "admin",
		Password: crypto.Password(pass),
		Email:    "admin@example.org",
		Role:     UserStatusAdmin,
	}).Error
	if err != nil {
		log.Log().WithError(err).Error("failed to add default user")
		return
	}

	log.Log().Infof("added default user, username: %s, email: %s, password: %s", "admin", "admin@example.org", pass)
}

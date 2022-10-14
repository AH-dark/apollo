package vo

import (
	"github.com/AH-dark/apollo/model"
	"github.com/samber/lo"
)

type SettingsListVO struct {
	Settings map[string]string `json:"settings" xml:"Settings"`
}

func BuildSettingsListVO(settings []*model.Setting) SettingsListVO {
	settingsListVO := SettingsListVO{
		Settings: lo.SliceToMap(settings, func(setting *model.Setting) (string, string) {
			return setting.Name, setting.Value
		}),
	}

	return settingsListVO
}

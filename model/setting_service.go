package model

import "gorm.io/gorm"

type SettingService interface {
	GetSetting(name string) (*Setting, error)
	GetSettingsByType(settingType SettingType) ([]*Setting, error)
	GetSettings([]string) (map[string]*Setting, error)

	Exist(name string) bool
	Save(setting *Setting) error
	SaveAll(settings []*Setting) error
}

type settingService struct {
	db *gorm.DB
}

func NewSettingService(db *gorm.DB) SettingService {
	return &settingService{db: db}
}

func (s *settingService) GetSetting(name string) (*Setting, error) {
	var setting Setting
	err := s.db.Where("name = ?", name).First(&setting).Error
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

func (s *settingService) GetSettingsByType(settingType SettingType) ([]*Setting, error) {
	var settings []*Setting
	err := s.db.Where("type = ?", settingType).Find(&settings).Error
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (s *settingService) GetSettings(names []string) (map[string]*Setting, error) {
	var settings []*Setting
	err := s.db.Where("name in (?)", names).Find(&settings).Error
	if err != nil {
		return nil, err
	}
	m := make(map[string]*Setting)
	for _, setting := range settings {
		m[setting.Name] = setting
	}
	return m, nil
}

func (s *settingService) Exist(name string) bool {
	var setting Setting
	err := s.db.Where("name = ?", name).First(&setting).Error
	return err == nil
}

func (s *settingService) Save(setting *Setting) error {
	return s.db.Save(setting).Error
}

func (s *settingService) SaveAll(settings []*Setting) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, setting := range settings {
			if err := tx.Save(setting).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

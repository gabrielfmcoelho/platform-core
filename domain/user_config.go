package domain

import (
	"gorm.io/gorm"
)

type UserConfig struct {
	gorm.Model
	UserID          uint                `gorm:"not null"`
	ServicesConfigs []UserServiceConfig `gorm:"foreignKey:UserID"`
}

type UserConfigRepository interface {
	Create(userConfig *UserConfig) error
	Fetch() ([]UserConfig, error)
	GetByID(id uint) (UserConfig, error)
	GetByUserID(userID uint) (UserConfig, error)
	Update(userConfigID uint, userConfig *UserConfig) error
	Delete(userConfigID uint) error
}

type UserConfigUsecase interface {
	Create(userConfig *UserConfig) error
	Fetch() ([]UserConfig, error)
	GetByID(id uint) (UserConfig, error)
	GetByUserID(userID uint) (UserConfig, error)
	Update(userConfigID uint, userConfig *UserConfig) error
	Delete(userConfigID uint) error
}

package domain

import (
	"context"

	"gorm.io/gorm"
)

type UserServiceConfig struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ServiceID uint `gorm:"not null"`
	IsPinned  bool `gorm:"default:false"`
}

type UserServiceConfigRepository interface {
	Create(ctx context.Context, userServiceConfig *UserServiceConfig) error
	Fetch(ctx context.Context) ([]UserServiceConfig, error)
	GetByID(ctx context.Context, id uint) (UserServiceConfig, error)
	GetByUserID(ctx context.Context, userID uint) (UserServiceConfig, error)
	GetByServiceID(ctx context.Context, serviceID uint) (UserServiceConfig, error)
	Update(ctx context.Context, userServiceConfigID uint, userServiceConfig *UserServiceConfig) error
	Delete(ctx context.Context, userServiceConfigID uint) error
}

type UserServiceConfigUsecase interface {
	Create(ctx context.Context, userServiceConfig *UserServiceConfig) error
	Fetch(ctx context.Context) ([]UserServiceConfig, error)
	GetByID(ctx context.Context, id uint) (UserServiceConfig, error)
	GetByUserID(ctx context.Context, userID uint) (UserServiceConfig, error)
	GetByServiceID(ctx context.Context, serviceID uint) (UserServiceConfig, error)
	Update(ctx context.Context, userServiceConfigID uint, userServiceConfig *UserServiceConfig) error
	Delete(ctx context.Context, userServiceConfigID uint) error
}

package domain

import (
	"context"

	"gorm.io/gorm"
)

type UserServiceLog struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	ServiceID uint   `gorm:"not null"`
	Date      string `gorm:"size:255"`
	Duration  string `gorm:"size:255"`
}

type UserServiceLogRepository interface {
	Create(ctx context.Context, UserServiceLog *UserServiceLog) error
	Fetch(ctx context.Context) ([]UserServiceLog, error)
	GetByID(ctx context.Context, id uint) (UserServiceLog, error)
	GetByUserID(ctx context.Context, userID uint) (UserServiceLog, error)
	GetByServiceID(ctx context.Context, serviceID uint) (UserServiceLog, error)
	Update(ctx context.Context, UserServiceLogID uint, UserServiceLog *UserServiceLog) error
	Delete(ctx context.Context, UserServiceLogID uint) error
}

type UserServiceLogUsecase interface {
	Create(ctx context.Context, UserServiceLog *UserServiceLog) error
	Fetch(ctx context.Context) ([]UserServiceLog, error)
	GetByID(ctx context.Context, id uint) (UserServiceLog, error)
	GetByUserID(ctx context.Context, userID uint) (UserServiceLog, error)
	GetByServiceID(ctx context.Context, serviceID uint) (UserServiceLog, error)
	Update(ctx context.Context, UserServiceLogID uint, UserServiceLog *UserServiceLog) error
	Delete(ctx context.Context, UserServiceLogID uint) error
}

package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// MANY TO ONE WITH USER

type UserLog struct {
	gorm.Model
	UserID    uint   `gorm:"not null;Index"`
	IPAddress string `gorm:"size:255"`
	Action    string `gorm:"size:255;not null"`
}

type UserLogRepository interface {
	Create(ctx context.Context, userLog *UserLog) error
	Fetch(ctx context.Context) ([]UserLog, error)
	GetByUserID(ctx context.Context, userID uint) ([]UserLog, error)
	GetByDate(ctx context.Context, userID uint, date time.Time) ([]UserLog, error)
	DeleteByID(ctx context.Context, userLogID uint) error
}

type UserLogUsecase interface {
	Create(ctx context.Context, userLog *UserLog) error
	Fetch(ctx context.Context) ([]UserLog, error)
	GetByID(ctx context.Context, id uint) (UserLog, error)
	GetByUserID(ctx context.Context, userID uint) (UserLog, error)
	Update(ctx context.Context, userLogID uint, userLog *UserLog) error
	Delete(ctx context.Context, userLogID uint) error
}

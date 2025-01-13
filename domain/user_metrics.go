package domain

import (
	"context"

	"gorm.io/gorm"
)

// ONE TO ONE WITH USER

type UserMetrics struct {
	gorm.Model
	UserID             uint   `gorm:"not null;uniqueIndex"`
	FavoriteServiceID  uint   `gorm:""`
	LastIP             string `gorm:"size:255"`
	LastLogin          string `gorm:"size:255"`
	TotalLogins        int    `gorm:"size:255"`
	TotalUsageDuration int    `gorm:"size:255"`
}

type UserMetricsRepository interface {
	Create(ctx context.Context, userMetrics *UserMetrics) error
	Fetch(ctx context.Context) ([]UserMetrics, error)
	GetByID(ctx context.Context, id uint) (UserMetrics, error)
	GetByUserID(ctx context.Context, userID uint) (UserMetrics, error)
	Update(ctx context.Context, userMetricsID uint, userMetrics *UserMetrics) error
	Delete(ctx context.Context, userMetricsID uint) error
}

type UserMetricsUsecase interface {
	Create(ctx context.Context, userMetrics *UserMetrics) error
	Fetch(ctx context.Context) ([]UserMetrics, error)
	GetByID(ctx context.Context, id uint) (UserMetrics, error)
	GetByUserID(ctx context.Context, userID uint) (UserMetrics, error)
	Update(ctx context.Context, userMetricsID uint, userMetrics *UserMetrics) error
	Delete(ctx context.Context, userMetricsID uint) error
}

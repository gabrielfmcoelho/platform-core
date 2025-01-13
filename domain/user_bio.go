package domain

import (
	"context"

	"gorm.io/gorm"
)

// ONE TO ONE WITH USER

type UserBio struct {
	gorm.Model
	UserID    uint   `gorm:"not null;uniqueIndex"`
	FirstName string `gorm:"size:255"`
	SurName   string `gorm:"size:255"`
	Position  string `gorm:"size:255"`
	Phone     string `gorm:"size:255"`
	Sex       string `gorm:"size:255"`
}

type UserBioRepository interface {
	Create(ctx context.Context, userBio *UserBio) error
	Fetch(ctx context.Context) ([]UserBio, error)
	GetByID(ctx context.Context, id uint) (UserBio, error)
	GetByUserID(ctx context.Context, userID uint) (UserBio, error)
	Update(ctx context.Context, userBioID uint, userBio *UserBio) error
	Delete(ctx context.Context, userBioID uint) error
}

type UserBioUsecase interface {
	Create(ctx context.Context, userBio *UserBio) error
	Fetch(ctx context.Context) ([]UserBio, error)
	GetByID(ctx context.Context, id uint) (UserBio, error)
	GetByUserID(ctx context.Context, userID uint) (UserBio, error)
	Update(ctx context.Context, userBioID uint, userBio *UserBio) error
	Delete(ctx context.Context, userBioID uint) error
}

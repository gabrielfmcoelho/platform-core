package domain

import (
	"context"

	"gorm.io/gorm"
)

// Admin, Manager, User, Guest

type UserRole struct {
	gorm.Model
	RoleName string `gorm:"size:255;uniqueIndex;not null"`
}

type UserRoleRepository interface {
	Create(ctx context.Context, userRole *UserRole) error
	Fetch(ctx context.Context) ([]UserRole, error)
	GetByID(ctx context.Context, id uint) (UserRole, error)
	GetByRoleName(ctx context.Context, roleName string) (UserRole, error)
	Update(ctx context.Context, userRoleID uint, userRole *UserRole) error
	Delete(ctx context.Context, userRoleID uint) error
}

type UserRoleUsecase interface {
	Create(ctx context.Context, userRole *UserRole) error
	Fetch(ctx context.Context) ([]UserRole, error)
	GetByIdentifier(ctx context.Context, identifier string) (UserRole, error)
	Update(ctx context.Context, userRoleID uint, userRole *UserRole) error
	Delete(ctx context.Context, userRoleID uint) error
}

package domain

import (
	"context"

	"gorm.io/gorm"
)

type OrganizationRole struct {
	gorm.Model
	RoleName string `gorm:"size:255;uniqueIndex;not null"`
}

type OrganizationRoleRepository interface {
	Create(ctx context.Context, organizationRole *OrganizationRole) error
	Fetch(ctx context.Context) ([]OrganizationRole, error)
	GetByID(ctx context.Context, id uint) (OrganizationRole, error)
	GetByRoleName(ctx context.Context, roleName string) (OrganizationRole, error)
	Update(ctx context.Context, organizationRoleID uint, organizationRole *OrganizationRole) error
	Delete(ctx context.Context, organizationRoleID uint) error
}

type OrganizationRoleUsecase interface {
	Create(ctx context.Context, organizationRole *OrganizationRole) error
	Fetch(ctx context.Context) ([]OrganizationRole, error)
	GetByIdentifier(ctx context.Context, identifier string) (OrganizationRole, error)
	Update(ctx context.Context, organizationRoleID uint, organizationRole *OrganizationRole) error
	Delete(ctx context.Context, organizationRoleID uint) error
}

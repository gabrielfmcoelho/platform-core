package domain

import (
	"context"

	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name               string                   `gorm:"size:255;uniqueIndex;not null"`
	Nickname           string                   `gorm:"size:255"`
	LogoUrl            string                   `gorm:"size:255"`
	RoleID             uint                     `gorm:"not null"`
	Role               OrganizationRole         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Users              []User                   `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subscription       OrganizationSubscription `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SubscribedServices []Service                `gorm:"many2many:organization_services;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Metrics            OrganizationMetrics      `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CreateOrganization struct {
	Name               string `json:"name" binding:"required"`
	OrganizationRoleID uint   `json:"organization_role_id" binding:"required,number"`
}

type PublicOrganization struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	LogoUrl  string `json:"logo_url"`
}

type OrganizationRepository interface {
	Create(ctx context.Context, organization *Organization) error
	Fetch(ctx context.Context) ([]Organization, error)
	GetByID(ctx context.Context, id uint) (Organization, error)
	GetByName(ctx context.Context, name string) (Organization, error)
	GetUsers(ctx context.Context, id uint) ([]User, error)
	GetSubscribedServices(ctx context.Context, id uint) ([]PublicService, error)
	Update(ctx context.Context, organizationID uint, organization *Organization) error
	Delete(ctx context.Context, organizationID uint) error
}

type OrganizationUsecase interface {
	Create(ctx context.Context, organization *Organization) error
	Fetch(ctx context.Context) ([]PublicOrganization, error)
	GetByIdentifier(ctx context.Context, identifier string) (PublicOrganization, error)
	GetUsers(ctx context.Context, id uint) ([]PublicUser, error)
	GetSubscribedServices(ctx context.Context, id uint) ([]PublicService, error)
	Update(ctx context.Context, organizationID uint, organization *Organization) error
	Delete(ctx context.Context, organizationID uint) error
}

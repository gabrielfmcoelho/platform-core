package domain

import (
	"context"

	"gorm.io/gorm"
)

// ONE TO ONE WITH ORGANIZATION

type OrganizationSubscription struct {
	gorm.Model
	Active                   bool    `gorm:"default:false"`
	OrganizationID           uint    `gorm:"not null;uniqueIndex"`
	SubscriptionValue        float64 `gorm:"not null"`
	SubscriptionPeriod       string  `gorm:"not null"`
	SubscriptionUsersLimit   int     `gorm:"not null"`
	SubscriptionReportsLimit int     `gorm:"not null"`
	SubscriptionInitDate     string  `gorm:"not null"`
	SubscriptionEndDate      string  `gorm:"not null"`
}

type OrganizationSubscriptionRepository interface {
	Create(ctx context.Context, organizationSubscription *OrganizationSubscription) error
	Fetch(ctx context.Context) ([]OrganizationSubscription, error)
	GetByID(ctx context.Context, id uint) (OrganizationSubscription, error)
	GetByOrganizationID(ctx context.Context, organizationID uint) (OrganizationSubscription, error)
	Update(ctx context.Context, organizationSubscriptionID uint, organizationSubscription *OrganizationSubscription) error
	Delete(ctx context.Context, organizationSubscriptionID uint) error
}

type OrganizationSubscriptionUsecase interface {
	Create(ctx context.Context, organizationSubscription *OrganizationSubscription) error
	Fetch(ctx context.Context) ([]OrganizationSubscription, error)
	GetByID(ctx context.Context, id uint) (OrganizationSubscription, error)
	GetByOrganizationID(ctx context.Context, organizationID uint) (OrganizationSubscription, error)
	Update(ctx context.Context, organizationSubscriptionID uint, organizationSubscription *OrganizationSubscription) error
	Delete(ctx context.Context, organizationSubscriptionID uint) error
}

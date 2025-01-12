package domain

import (
	"context"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	MarketingName string         `gorm:"size:255;uniqueIndex;not null"`
	Name          string         `gorm:"size:255;uniqueIndex;not null"`
	Description   string         `gorm:"size:255;not null"`
	AppUrl        string         `gorm:"size:255;not null"`
	IconUrl       string         `gorm:"size:255"`
	TagLine       string         `gorm:"size:255"`
	Benefits      string         `gorm:"size:255"`
	Features      string         `gorm:"size:255"`
	Tags          string         `gorm:"size:255"`
	ScreenshotUrl string         `gorm:"size:255"`
	LastUpdate    string         `gorm:"size:255"`
	Status        string         `gorm:"size:255"`
	Price         float64        `gorm:"not null"`
	Organization  []Organization `gorm:"many2many:organization_services;"`
}

type PublicService struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ServiceRepository interface {
	Create(ctx context.Context, service *Service) error
	Fetch(ctx context.Context) ([]Service, error)
	GetByID(ctx context.Context, id uint) (Service, error)
	GetByName(ctx context.Context, name string) (Service, error)
	Update(ctx context.Context, serviceID uint, service *Service) error
	Delete(ctx context.Context, serviceID uint) error
}

type ServiceUsecase interface {
	Create(ctx context.Context, service *Service) error
	Fetch(ctx context.Context) ([]PublicService, error)
	GetByIdentifier(ctx context.Context, identifier string) (PublicService, error)
	Update(ctx context.Context, serviceID uint, service *Service) error
	Delete(ctx context.Context, serviceID uint) error
}

package domain

import (
	"context"

	"gorm.io/gorm"
)

// MANY TO MANY WITH ORGANIZATION

type Service struct {
	gorm.Model
	MarketingName string         `gorm:"size:255;uniqueIndex;not null"`
	Name          string         `gorm:"size:255;uniqueIndex;not null"`
	Description   string         `gorm:"size:255;not null"`
	AppUrl        string         `gorm:"size:255;not null"`
	IconUrl       string         `gorm:"size:255"`
	ScreenshotUrl string         `gorm:"size:255"`
	TagLine       string         `gorm:"size:255"`
	Benefits      string         `gorm:"size:255"`
	Features      string         `gorm:"size:255"`
	Tags          string         `gorm:"size:255"`
	LastUpdate    string         `gorm:"size:255"`
	Status        string         `gorm:"size:255"`
	Price         float64        `gorm:"not null"`
	IsMarketing   bool           `gorm:"not null;default:false"`
	Organization  []Organization `gorm:"many2many:organization_services;"`
}

type PublicService struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	AppUrl     string `json:"app_url"`
	LastUpdate string `json:"last_update"`
	Status     string `json:"status"`
}

type HubService struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	IconUrl       string  `json:"icon_url"`
	Description   string  `json:"description"`
	ScreenshotUrl string  `json:"screenshot_url"`
	LastUpdate    string  `json:"last_update"`
	Status        string  `json:"status"`
	Price         float64 `json:"price"`
}

type MarketingService struct {
	ID            uint     `json:"id"`
	IconUrl       string   `json:"icon_url"`
	MarketingName string   `json:"marketing_name"`
	TagLine       string   `json:"tag_line"`
	Description   string   `json:"description"`
	Benefits      []string `json:"benefits"`
	Features      []string `json:"features"`
	Tags          []string `json:"tags"`
}

type UseService struct {
	Service PublicService `json:"service"`
	LogID   uint          `json:"log_id"`
}

type Heartbeat struct {
	LogID    uint `json:"log_id"`
	Duration int  `json:"duration"`
}

type ServiceRepository interface {
	Create(ctx context.Context, service *Service) error
	Fetch(ctx context.Context) ([]Service, error)
	GetByID(ctx context.Context, id uint) (Service, error)
	GetByName(ctx context.Context, name string) (Service, error)
	GetByOrganization(ctx context.Context, organizationID uint) ([]Service, error)
	GetMarketing(ctx context.Context) ([]Service, error)
	SetAvailabilityToOrganization(ctx context.Context, serviceID uint, organizationID uint) error
	Update(ctx context.Context, serviceID uint, service *Service) error
	Delete(ctx context.Context, serviceID uint) error
}

type ServiceUsecase interface {
	Create(ctx context.Context, service *Service) error
	Fetch(ctx context.Context) ([]PublicService, error)
	GetByIdentifier(ctx context.Context, identifier string) (PublicService, error)
	GetByOrganization(ctx context.Context, organizationID uint) ([]HubService, error)
	GetMarketing(ctx context.Context) ([]MarketingService, error)
	SetAvailabilityToOrganization(ctx context.Context, serviceID uint, organizationID uint) error
	Use(ctx context.Context, userID uint, serviceID uint) (UseService, uint, error)
	Heartbeat(ctx context.Context, logID uint, duration int) error
	Update(ctx context.Context, serviceID uint, service *Service) error
	Delete(ctx context.Context, serviceID uint) error
}

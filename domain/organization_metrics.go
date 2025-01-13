package domain

import (
	"gorm.io/gorm"
)

// ONE TO ONE WITH ORGANIZATION

type OrganizationMetrics struct {
	gorm.Model
	OrganizationID           uint   `gorm:"not null;uniqueIndex"`
	TotalServices            int    `gorm:"default:0"`
	TotalUsers               int    `gorm:"default:0"`
	TotalReports             int    `gorm:"default:0"`
	TotalReportsCurrentMonth int    `gorm:"default:0"`
	LastReportDate           string `gorm:"size:255"`
	NextReportDate           string `gorm:"size:255"`
}

type OrganizationMetricsRepository interface {
	Create(organizationMetrics *OrganizationMetrics) error
	Fetch() ([]OrganizationMetrics, error)
	GetByID(id uint) (OrganizationMetrics, error)
	GetByOrganizationID(organizationID uint) (OrganizationMetrics, error)
	Update(organizationMetricsID uint, organizationMetrics *OrganizationMetrics) error
	Delete(organizationMetricsID uint) error
}

type OrganizationMetricsUsecase interface {
	Create(organizationMetrics *OrganizationMetrics) error
	Fetch() ([]OrganizationMetrics, error)
	GetByID(id uint) (OrganizationMetrics, error)
	GetByOrganizationID(organizationID uint) (OrganizationMetrics, error)
	Update(organizationMetricsID uint, organizationMetrics *OrganizationMetrics) error
	Delete(organizationMetricsID uint) error
}

package repository

import (
	"context"
	"errors"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

type organizationRepository struct {
	db *gorm.DB
}

// NewOrganizationRepository retorna uma instância que implementa a interface OrganizationRepository
func NewOrganizationRepository(db *gorm.DB) domain.OrganizationRepository {
	return &organizationRepository{db: db}
}

// Create cria uma nova Organização no banco de dados
func (r *organizationRepository) Create(ctx context.Context, organization *domain.Organization) error {
	if err := r.db.WithContext(ctx).Create(organization).Error; err != nil {
		return err
	}
	return nil
}

// Fetch retorna todas as Organizações do banco de dados
func (r *organizationRepository) Fetch(ctx context.Context) ([]domain.Organization, error) {
	var orgs []domain.Organization
	if err := r.db.WithContext(ctx).Preload("Role").
		Preload("Users").
		Preload("SubscribedServices").
		Find(&orgs).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}

// GetByID retorna uma Organização específica baseada no ID
func (r *organizationRepository) GetByID(ctx context.Context, id uint) (domain.Organization, error) {
	var org domain.Organization
	if err := r.db.WithContext(ctx).
		Preload("Role").
		Preload("Users").
		Preload("SubscribedServices").
		First(&org, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return org, domain.ErrNotFound
		}
		return org, err
	}
	return org, nil
}

// GetByName retorna uma Organização específica baseada no nome
func (r *organizationRepository) GetByName(ctx context.Context, name string) (domain.Organization, error) {
	var org domain.Organization
	if err := r.db.WithContext(ctx).
		Preload("Role").
		Preload("Users").
		Preload("SubscribedServices").
		Where("name = ?", name).
		First(&org).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return org, domain.ErrNotFound
		}
		return org, err
	}
	return org, nil
}

// GetUsers retorna os usuários de uma Organização
func (r *organizationRepository) GetUsers(ctx context.Context, organizationID uint) ([]domain.User, error) {
	var org domain.Organization
	if err := r.db.WithContext(ctx).
		Preload("Users").
		First(&org, organizationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return org.Users, nil
}

// GetSubscribedServices retorna os serviços que a Organização está inscrita (many2many)
func (r *organizationRepository) GetSubscribedServices(ctx context.Context, organizationID uint) ([]domain.PublicService, error) {
	var org domain.Organization
	if err := r.db.WithContext(ctx).
		Preload("SubscribedServices").
		First(&org, organizationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	// Converter []Service para []PublicService (supondo que sua domain tenha mapeamento)
	var publicServices []domain.PublicService
	for _, srv := range org.SubscribedServices {
		publicServices = append(publicServices, domain.PublicService{
			ID:   srv.ID,
			Name: srv.Name,
		})
	}
	return publicServices, nil
}

// Update atualiza dados de uma Organização
func (r *organizationRepository) Update(ctx context.Context, organizationID uint, data *domain.Organization) error {
	// Checa se existe a org
	if err := r.db.WithContext(ctx).First(&domain.Organization{}, organizationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrNotFound
		}
		return err
	}

	// Atualiza
	if err := r.db.WithContext(ctx).
		Model(&domain.Organization{}).
		Where("id = ?", organizationID).
		Updates(data).
		Error; err != nil {
		return err
	}
	return nil
}

// Delete remove (fisicamente) uma Organização
func (r *organizationRepository) Delete(ctx context.Context, organizationID uint) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Organization{}, organizationID).Error; err != nil {
		return err
	}
	return nil
}

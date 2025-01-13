package repository

import (
	"context"
	"errors"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

type organizationRoleRepository struct {
	db *gorm.DB
}

// NewOrganizationRoleRepository retorna uma instância que implementa a interface OrganizationRoleRepository
func NewOrganizationRoleRepository(db *gorm.DB) domain.OrganizationRoleRepository {
	return &organizationRoleRepository{
		db: db,
	}
}

// Create cria um novo OrganizationRole no banco
func (r *organizationRoleRepository) Create(ctx context.Context, orgRole *domain.OrganizationRole) error {
	if err := r.db.WithContext(ctx).Create(orgRole).Error; err != nil {
		return err
	}
	return nil
}

// Fetch retorna todos os OrganizationRoles
func (r *organizationRoleRepository) Fetch(ctx context.Context) ([]domain.OrganizationRole, error) {
	var roles []domain.OrganizationRole
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// GetByID retorna um OrganizationRole específico pelo ID
func (r *organizationRoleRepository) GetByID(ctx context.Context, id uint) (domain.OrganizationRole, error) {
	var role domain.OrganizationRole
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return role, domain.ErrNotFound
		}
		return role, err
	}
	return role, nil
}

// GetByRoleName retorna um OrganizationRole específico pelo nome
func (r *organizationRoleRepository) GetByRoleName(ctx context.Context, roleName string) (domain.OrganizationRole, error) {
	var role domain.OrganizationRole
	if err := r.db.WithContext(ctx).Where("role_name = ?", roleName).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return role, domain.ErrNotFound
		}
		return role, err
	}
	return role, nil
}

// Update atualiza um OrganizationRole
func (r *organizationRoleRepository) Update(ctx context.Context, orgRoleID uint, updated *domain.OrganizationRole) error {
	if err := r.db.WithContext(ctx).
		Model(&domain.OrganizationRole{}).
		Where("id = ?", orgRoleID).
		Updates(updated).
		Error; err != nil {
		return err
	}
	return nil
}

// Delete remove (fisicamente) um OrganizationRole
func (r *organizationRoleRepository) Delete(ctx context.Context, orgRoleID uint) error {
	if err := r.db.WithContext(ctx).Delete(&domain.OrganizationRole{}, orgRoleID).Error; err != nil {
		return err
	}
	return nil
}

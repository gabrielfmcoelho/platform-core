package repository

import (
	"context"
	"errors"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

type userRoleRepository struct {
	db *gorm.DB
}

// NewUserRoleRepository retorna uma instância que implementa a interface UserRoleRepository
func NewUserRoleRepository(db *gorm.DB) domain.UserRoleRepository {
	return &userRoleRepository{
		db: db,
	}
}

// Create cria um novo UserRole no banco
func (r *userRoleRepository) Create(ctx context.Context, userRole *domain.UserRole) error {
	if err := r.db.WithContext(ctx).Create(userRole).Error; err != nil {
		return err
	}
	return nil
}

// Fetch retorna todos os UserRoles
func (r *userRoleRepository) Fetch(ctx context.Context) ([]domain.UserRole, error) {
	var roles []domain.UserRole
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// GetByID retorna um UserRole específico pelo ID
func (r *userRoleRepository) GetByID(ctx context.Context, id uint) (domain.UserRole, error) {
	var role domain.UserRole
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return role, domain.ErrNotFound
		}
		return role, err
	}
	return role, nil
}

// GetByRoleName retorna um UserRole específico pelo nome
func (r *userRoleRepository) GetByRoleName(ctx context.Context, roleName string) (domain.UserRole, error) {
	var role domain.UserRole
	if err := r.db.WithContext(ctx).Where("role_name = ?", roleName).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return role, domain.ErrNotFound
		}
		return role, err
	}
	return role, nil
}

// Update atualiza um UserRole
func (r *userRoleRepository) Update(ctx context.Context, userRoleID uint, updated *domain.UserRole) error {
	if err := r.db.WithContext(ctx).
		Model(&domain.UserRole{}).
		Where("id = ?", userRoleID).
		Updates(updated).
		Error; err != nil {
		return err
	}
	return nil
}

// Delete remove (fisicamente) um UserRole
func (r *userRoleRepository) Delete(ctx context.Context, userRoleID uint) error {
	if err := r.db.WithContext(ctx).Delete(&domain.UserRole{}, userRoleID).Error; err != nil {
		return err
	}
	return nil
}

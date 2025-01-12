package repository

import (
	"context"
	"errors"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create cria um novo usuário no banco de dados
func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	// Usando a transação, se necessário
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return domain.ErrDataBaseInternalError
	}
	return nil
}

// Fetch retorna todos os usuários do banco de dados
func (r *userRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, domain.ErrDataBaseInternalError
	}
	return users, nil
}

// GetByEmail retorna um usuário específico com base no email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, domain.ErrUserEmailNotFound
		}
		return user, domain.ErrDataBaseInternalError
	}
	return user, nil
}

// GetByID retorna um usuário específico com base no ID
func (r *userRepository) GetByID(ctx context.Context, id uint) (domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, domain.ErrNotFound
		}
		return user, domain.ErrDataBaseInternalError
	}
	return user, nil
}

// Update atualiza os dados de um usuário no banco
func (r *userRepository) Update(ctx context.Context, userID uint, userData *domain.User) error {
	// check if the user exists and get the user
	if err := r.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("id = ?", userID).
		Updates(userData).Error; err != nil {
		return domain.ErrDataBaseInternalError
	}
	return nil
}

// Archive marca um usuário como arquivado (soft delete)
func (r *userRepository) Archive(ctx context.Context, userID uint) error {
	// Alterando apenas o campo IsArchived
	// get user by ID
	user, err := r.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	user.IsArchived = true
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return domain.ErrDataBaseInternalError
	}
	return nil
}

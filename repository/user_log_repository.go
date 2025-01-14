package repository

import (
	"context"

	"time"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

type userLogRepository struct {
	db *gorm.DB
}

func NewUserLogRepository(db *gorm.DB) domain.UserLogRepository {
	return &userLogRepository{
		db: db,
	}
}

// Create cria um novo log de usuário no banco
func (r *userLogRepository) Create(ctx context.Context, userLog *domain.UserLog) error {
	if err := r.db.WithContext(ctx).Create(userLog).Error; err != nil {
		return err
	}
	return nil
}

// Fetch retorna todos os logs de usuário
func (r *userLogRepository) Fetch(ctx context.Context) ([]domain.UserLog, error) {
	var logs []domain.UserLog
	if err := r.db.WithContext(ctx).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

// GetByUserID retorna todos os logs de um usuário específico
func (r *userLogRepository) GetByUserID(ctx context.Context, userID uint) ([]domain.UserLog, error) {
	var logs []domain.UserLog
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

// GetByDate retorna todos os logs de um usuário específico em uma data específica
func (r *userLogRepository) GetByDate(ctx context.Context, userID uint, date time.Time) ([]domain.UserLog, error) {
	var logs []domain.UserLog
	if err := r.db.WithContext(ctx).Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, date, date.AddDate(0, 0, 1)).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

// DeleteByID deleta um log de usuário específico pelo ID
func (r *userLogRepository) DeleteByID(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&domain.UserLog{}, id).Error; err != nil {
		return err
	}
	return nil
}

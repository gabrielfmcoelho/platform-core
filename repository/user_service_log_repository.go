package repository

import (
	"context"
	"errors"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

type userServiceLogRepository struct {
	db *gorm.DB
}

func NewUserServiceLogRepository(db *gorm.DB) domain.UserServiceLogRepository {
	return &userServiceLogRepository{
		db: db,
	}
}

// Create inserts a new UserServiceLog in the database
func (r *userServiceLogRepository) Create(ctx context.Context, userServiceLog *domain.UserServiceLog) error {
	if err := r.db.WithContext(ctx).Create(userServiceLog).Error; err != nil {
		// Adjust to your error handling
		return domain.ErrDataBaseInternalError
	}
	return nil
}

// Fetch returns all UserServiceLog entries
func (r *userServiceLogRepository) Fetch(ctx context.Context) ([]domain.UserServiceLog, error) {
	var logs []domain.UserServiceLog
	if err := r.db.WithContext(ctx).Find(&logs).Error; err != nil {
		return nil, domain.ErrDataBaseInternalError
	}
	return logs, nil
}

// GetByID returns a UserServiceLog by its ID
func (r *userServiceLogRepository) GetByID(ctx context.Context, id uint) (domain.UserServiceLog, error) {
	var log domain.UserServiceLog
	if err := r.db.WithContext(ctx).First(&log, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return log, domain.ErrNotFound
		}
		return log, domain.ErrDataBaseInternalError
	}
	return log, nil
}

// GetByUserID returns a UserServiceLog by user ID
func (r *userServiceLogRepository) GetByUserID(ctx context.Context, userID uint) (domain.UserServiceLog, error) {
	var log domain.UserServiceLog
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&log).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return log, domain.ErrNotFound
		}
		return log, domain.ErrDataBaseInternalError
	}
	return log, nil
}

// GetByServiceID returns a UserServiceLog by service ID
func (r *userServiceLogRepository) GetByServiceID(ctx context.Context, serviceID uint) (domain.UserServiceLog, error) {
	var log domain.UserServiceLog
	if err := r.db.WithContext(ctx).Where("service_id = ?", serviceID).First(&log).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return log, domain.ErrNotFound
		}
		return log, domain.ErrDataBaseInternalError
	}
	return log, nil
}

// Update updates an existing UserServiceLog, add duration to the existing duration
func (r *userServiceLogRepository) UpdateDuration(ctx context.Context, userServiceLogID uint, duration int) error {
	if err := r.db.WithContext(ctx).Model(&domain.UserServiceLog{}).
		Where("id = ?", userServiceLogID).
		Update("duration", gorm.Expr("duration + ?", duration)).Error; err != nil {
		return domain.ErrDataBaseInternalError
	}
	return nil
}

// Delete removes a UserServiceLog by its ID (hard delete)
func (r *userServiceLogRepository) Delete(ctx context.Context, userServiceLogID uint) error {
	if err := r.db.WithContext(ctx).Delete(&domain.UserServiceLog{}, userServiceLogID).Error; err != nil {
		return domain.ErrDataBaseInternalError
	}
	return nil
}

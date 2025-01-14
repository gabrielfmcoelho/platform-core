package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"github.com/gabrielfmcoelho/platform-core/internal"
)

type userServiceLogUsecase struct {
	userServiceLogRepo domain.UserServiceLogRepository
	contextTimeout     time.Duration
}

func NewUserServiceLogUsecase(
	repo domain.UserServiceLogRepository,
	timeout time.Duration,
) domain.UserServiceLogUsecase {
	return &userServiceLogUsecase{
		userServiceLogRepo: repo,
		contextTimeout:     timeout,
	}
}

// Fetch all UserServiceLog entries
func (u *userServiceLogUsecase) Fetch(ctx context.Context) ([]domain.UserServiceLog, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	logs, err := u.userServiceLogRepo.Fetch(c)
	if err != nil {
		if errors.Is(err, domain.ErrDataBaseInternalError) {
			return nil, domain.ErrDataBaseInternalError
		}
		return nil, domain.ErrInternalServerError
	}
	return logs, nil
}

// GetByIdentifier tries to parse identifier to either:
// - a numeric ID -> GetByID
// - or if it starts with "user:" -> parse the rest as userID
// - or if it starts with "service:" -> parse the rest as serviceID
// Otherwise returns ErrInvalidIdentifier
func (u *userServiceLogUsecase) GetByIdentifier(ctx context.Context, identifier string) (domain.UserServiceLog, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var log domain.UserServiceLog
	var err error

	// Example rules:
	// 1) if it's numeric: interpret as ID
	// 2) if starts with "user:" => parse userID
	// 3) if starts with "service:" => parse serviceID
	// else => err

	switch {
	case internal.IsNumeric(identifier):
		id, convErr := internal.ParseUint(identifier)
		if convErr != nil {
			return log, domain.ErrInvalidIdentifier
		}
		log, err = u.userServiceLogRepo.GetByID(c, id)

	case len(identifier) > 5 && identifier[:5] == "user:":
		userIDStr := identifier[5:]
		id, convErr := internal.ParseUint(userIDStr)
		if convErr != nil {
			return log, domain.ErrInvalidIdentifier
		}
		log, err = u.userServiceLogRepo.GetByUserID(c, id)

	case len(identifier) > 8 && identifier[:8] == "service:":
		serviceIDStr := identifier[8:]
		id, convErr := internal.ParseUint(serviceIDStr)
		if convErr != nil {
			return log, domain.ErrInvalidIdentifier
		}
		log, err = u.userServiceLogRepo.GetByServiceID(c, id)

	default:
		// not numeric, not user/service pattern
		return log, domain.ErrInvalidIdentifier
	}

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return log, domain.ErrNotFound
		}
		return log, domain.ErrInternalServerError
	}

	return log, nil
}

// Delete removes a UserServiceLog by ID
func (u *userServiceLogUsecase) Delete(ctx context.Context, userServiceLogID uint) error {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if err := u.userServiceLogRepo.Delete(c, userServiceLogID); err != nil {
		if errors.Is(err, domain.ErrDataBaseInternalError) {
			return domain.ErrDataBaseInternalError
		}
		return domain.ErrInternalServerError
	}
	return nil
}

package parser

import (
	"github.com/gabrielfmcoelho/platform-core/domain"
	"github.com/gabrielfmcoelho/platform-core/internal"
)

// Parse UserServiceLog to PublicUserServiceLog
func ToPublicUserServiceLog(log domain.UserServiceLog) domain.PublicUserServiceLog {
	return domain.PublicUserServiceLog{
		ID:        log.ID,
		UserID:    log.UserID,
		ServiceID: log.ServiceID,
		Duration:  internal.ToSeconds(log.Duration),
	}
}

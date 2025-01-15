package parser

import (
	"github.com/gabrielfmcoelho/platform-core/domain"
)

// Parse Any type struct to SuccessResponse
func ToSuccessResponse(data interface{}) domain.SuccessResponse {
	return domain.SuccessResponse{
		Message: "success",
		Data:    data,
	}
}

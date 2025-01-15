package controller

import (
	"net/http"
	"strconv"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"github.com/gabrielfmcoelho/platform-core/internal/parser"
	"github.com/gin-gonic/gin"
)

type UserServiceLogController struct {
	UserServiceLogUsecase domain.UserServiceLogUsecase
}

// FetchUserServiceLogs
// @Summary Fetch all UserServiceLogs
// @Description Gets all user-service log entries
// @Tags UserServiceLog
// @Produce json
// @Success 200 {array} domain.SuccessResponse{data=[]domain.PublicUserServiceLog}
// @Failure 500 {object} domain.ErrorResponse
// @Router /user-service-logs [get]
func (ctrl *UserServiceLogController) FetchUserServiceLogs(c *gin.Context) {
	logs, err := ctrl.UserServiceLogUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, parser.ToSuccessResponse(logs))
}

// GetUserServiceLogByIdentifier
// @Summary Get a UserServiceLog by ID or user:XXX or service:XXX
// @Description Gets a user-service log entry by a flexible identifier
// @Tags UserServiceLog
// @Produce json
// @Param identifier path string true "Identifier"
// @Success 200 {object} domain.SuccessResponse{data=domain.PublicUserServiceLog}
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /user-service-logs/{identifier} [get]
func (ctrl *UserServiceLogController) GetUserServiceLogByIdentifier(c *gin.Context) {
	identifier := c.Param("identifier")

	log, err := ctrl.UserServiceLogUsecase.GetByIdentifier(c, identifier)
	if err != nil {
		switch err {
		case domain.ErrInvalidIdentifier:
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		case domain.ErrNotFound:
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, parser.ToSuccessResponse(log))
}

// DeleteUserServiceLog
// @Summary Delete a UserServiceLog
// @Description Removes a user-service log record
// @Tags UserServiceLog
// @Produce json
// @Param logID path int true "UserServiceLog ID"
// @Success 204 "No Content"
// @Failure 500 {object} domain.ErrorResponse
// @Router /user-service-logs/{logID} [delete]
func (ctrl *UserServiceLogController) DeleteUserServiceLog(c *gin.Context) {
	logIDStr := c.Param("logID")
	logID, err := strconv.ParseUint(logIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid logID"})
		return
	}

	if err := ctrl.UserServiceLogUsecase.Delete(c, uint(logID)); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

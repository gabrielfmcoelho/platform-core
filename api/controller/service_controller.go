package controller

import (
	"fmt"
	"net/http"

	"strings"

	"github.com/gabrielfmcoelho/platform-core/bootstrap"
	"github.com/gabrielfmcoelho/platform-core/domain"
	"github.com/gabrielfmcoelho/platform-core/internal"
	"github.com/gabrielfmcoelho/platform-core/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	ServiceUsecase domain.ServiceUsecase
	Env            *bootstrap.Env
}

// CreateService cria um novo serviço
// @Summary Create Service
// @Description Creates a new service
// @Tags Service
// @Accept json
// @Produce json
// @Param service body domain.Service true "Service data"
// @Success 201 {object} domain.Service
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /services [post]
func (sc *ServiceController) CreateService(c *gin.Context) {
	var service domain.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := sc.ServiceUsecase.Create(c, &service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Retornar o service criado (ou alguma versão pública dele)
	c.JSON(http.StatusCreated, service)
}

// FetchServices retorna todos os serviços
// @Summary Fetch Services
// @Description Gets all available services
// @Tags Service
// @Produce json
// @Success 200 {array} domain.PublicService
// @Failure 500 {object} domain.ErrorResponse
// @Router /services [get]
func (sc *ServiceController) FetchServices(c *gin.Context) {
	services, err := sc.ServiceUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// GetServiceByIdentifier retorna um serviço por ID ou nome
// @Summary Get Service by Identifier
// @Description Gets service by numeric ID (e.g., /services/123) or name (/services/my-service)
// @Tags Service
// @Produce json
// @Param identifier path string true "Service ID or Name"
// @Success 200 {object} domain.PublicService
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /services/{identifier} [get]
func (sc *ServiceController) GetServiceByIdentifier(c *gin.Context) {
	identifier := c.Param("identifier")

	service, err := sc.ServiceUsecase.GetByIdentifier(c, identifier)
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, service)
}

// SetServiceAvailabilityToOrganization vincula um service a uma organização
// @Summary Set Service Availability
// @Description Links the service to an organization
// @Tags Service
// @Produce json
// @Param serviceID path int true "Service ID"
// @Param organizationID path int true "Organization ID"
// @Success 200 {object} domain.SuccessResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /services/{serviceID}/organization/{organizationID} [post]
func (sc *ServiceController) SetServiceAvailabilityToOrganization(c *gin.Context) {
	serviceID := c.Param("serviceID")
	organizationID := c.Param("organizationID")

	// converter para uint
	// (poderíamos extrair essa lógica para uma função utilitária)
	var sID, oID uint
	_, err := fmt.Sscanf(serviceID, "%d", &sID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid serviceID"})
		return
	}
	_, err = fmt.Sscanf(organizationID, "%d", &oID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid organizationID"})
		return
	}

	err = sc.ServiceUsecase.SetAvailabilityToOrganization(c, sID, oID)
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Service availability set successfully."})
}

// UseService
// @Summary Start using a service (create a usage log)
// @Description Logs that a user started using a service, returns log ID and public service data
// @Tags Service
// @Accept json
// @Produce json
// @Param serviceID path int true "Service ID"
// @Success 200 {object} domain.UseService
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /services/{serviceID}/use [post]
func (sc *ServiceController) UseService(c *gin.Context) {
	// 1) Parse serviceID from path
	serviceIDParam := c.Param("serviceID")
	var sID uint
	if _, err := fmt.Sscanf(serviceIDParam, "%d", &sID); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid serviceID"})
		return
	}

	// 2) extract UserID from jwt token
	authHeader := c.Request.Header.Get("Authorization")
	t := strings.Split(authHeader, " ")
	authToken := t[1]
	userID, err := tokenutil.ExtractIDFromToken(authToken, sc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
	}
	uID, err := internal.ParseUint(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid userID"})
		return
	}

	// 3) Call usecase
	service, logID, err := sc.ServiceUsecase.Use(c, uID, sID)
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		}
		return
	}

	// 4) Return result
	c.JSON(http.StatusOK, domain.UseService{
		LogID:   logID,
		Service: service,
	})
}

// HeartbeatService
// @Summary Heartbeat usage
// @Description Adds usage duration (in seconds) to a log record
// @Tags Service
// @Accept json
// @Produce json
// @Param heartbeat body domain.Heartbeat true "Heartbeat data"
// @Success 200 {object} domain.SuccessResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /services/heartbeat [patch]
func (sc *ServiceController) HeartbeatService(c *gin.Context) {
	// 1) Parse JSON body for logID and duration
	var req domain.Heartbeat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := sc.ServiceUsecase.Heartbeat(c, req.LogID, req.Duration)
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		}
		return
	}

	// 3) Return success
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Usage duration updated successfully"})
}

// UpdateService atualiza um service
// @Summary Update Service
// @Description Updates service data
// @Tags Service
// @Accept json
// @Produce json
// @Param serviceID path int true "Service ID"
// @Param service body domain.Service true "Service data"
// @Success 200 {object} domain.Service
// @Failure 500 {object} domain.ErrorResponse
// @Router /services/{serviceID} [put]
func (sc *ServiceController) UpdateService(c *gin.Context) {
	serviceID := c.Param("serviceID")

	var sID uint
	_, err := fmt.Sscanf(serviceID, "%d", &sID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid serviceID"})
		return
	}

	var service domain.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = sc.ServiceUsecase.Update(c, sID, &service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, service)
}

// DeleteService deleta um service
// @Summary Delete Service
// @Description Deletes a service
// @Tags Service
// @Produce json
// @Param serviceID path int true "Service ID"
// @Success 204 "No Content"
// @Failure 500 {object} domain.ErrorResponse
// @Router /services/{serviceID} [delete]
func (sc *ServiceController) DeleteService(c *gin.Context) {
	serviceID := c.Param("serviceID")

	var sID uint
	_, err := fmt.Sscanf(serviceID, "%d", &sID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid serviceID"})
		return
	}

	err = sc.ServiceUsecase.Delete(c, sID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Retornamos 204 No Content pois não há conteúdo no response
	c.Status(http.StatusNoContent)
}

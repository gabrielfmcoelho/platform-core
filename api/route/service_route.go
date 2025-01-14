package route

import (
	"time"

	"github.com/gabrielfmcoelho/platform-core/api/controller"
	"github.com/gabrielfmcoelho/platform-core/bootstrap"
	"github.com/gabrielfmcoelho/platform-core/repository"
	"github.com/gabrielfmcoelho/platform-core/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewServiceRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	sr := repository.NewServiceRepository(db)
	uslr := repository.NewUserServiceLogRepository(db)
	sc := &controller.ServiceController{
		ServiceUsecase: usecase.NewServiceUsecase(sr, uslr, timeout),
		Env:            env,
	}

	group.POST("/services", sc.CreateService)
	group.GET("/services", sc.FetchServices)
	group.GET("/services/:identifier", sc.GetServiceByIdentifier)
	group.POST("/services/:serviceID/organization/:organizationID", sc.SetServiceAvailabilityToOrganization)
	group.PUT("/services/:serviceID", sc.UpdateService)
	group.DELETE("/services/:serviceID", sc.DeleteService)
	group.POST("/services/:serviceID/use", sc.UseService)   // "start" usage
	group.PATCH("/services/heartbeat", sc.HeartbeatService) // "update" usage duration
}

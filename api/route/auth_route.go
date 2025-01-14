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

func NewAuthRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	ulr := repository.NewUserLogRepository(db)
	ac := &controller.AuthController{
		AuthUsecase: usecase.NewAuthUsecase(ur, ulr, timeout),
		Env:         env,
	}

	group.POST("/login", ac.Login)
	group.POST("/login-guest", ac.LoginGuest)
	group.POST("/forgot-password", ac.ForgotPassword)
	group.POST("/reset-password", ac.ResetPassword)
	group.POST("/refresh-token", ac.RefreshToken)
}

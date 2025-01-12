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
	lc := &controller.AuthController{
		AuthUsecase: usecase.NewAuthUsecase(ur, timeout),
		Env:         env,
	}

	group.POST("/login", lc.Login)
	group.POST("/forgot-password", lc.ForgotPassword)
	group.POST("/reset-password", lc.ResetPassword)
	group.POST("/refresh-token", lc.RefreshToken)
}

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

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	uc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		Env:         env,
	}

	group.POST("/user/create", uc.CreateUser)  // Create a new user account
	group.GET("/users", uc.FetchUsers)         // Get all users
	group.GET("/user/:identifier", uc.GetUser) // Get user by ID or email
	group.PUT("/user/:id", uc.UpdateUser)      // Update basic user information (email, password, etc)
	group.DELETE("/user/:id", uc.DeleteUser)   // Soft delete user (archive)
}

// DOUBT: How to implement query parameters in the routes in go?

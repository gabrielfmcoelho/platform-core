package controller

import (
	"net/http"

	"github.com/gabrielfmcoelho/platform-core/bootstrap"
	"github.com/gabrielfmcoelho/platform-core/domain"
	"github.com/gabrielfmcoelho/platform-core/internal"
	"github.com/gabrielfmcoelho/platform-core/internal/parser"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

// @Summary Create a new user account
// @Description Create a new user account with the input payload
// @Tags User
// @ID createUser
// @Accept json
// @Produce json
// @Param user body domain.CreateUser true "User object"
// @Success 201 "User created successfully"
// @Failure 400 {object} domain.ErrorResponse "Bad Request - Invalid input"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Router /user/create [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var user domain.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	err := uc.UserUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to create user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// @Summary Get all users
// @Description Get all users from the database
// @Tags User
// @ID fetchUsers
// @Produce json
// @Success 200 {object} domain.SuccessResponse{data=[]domain.PublicUser} "List of users"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Router /users [get]
func (uc *UserController) FetchUsers(c *gin.Context) {
	users, err := uc.UserUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to fetch users: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, parser.ToSuccessResponse(users))
}

// @Summary Get user by ID or email
// @Description Get user by ID or email
// @Tags User
// @ID getUser
// @Produce json
// @Param identifier path string true "User ID or email"
// @Success 200 {object} domain.SuccessResponse{data=domain.PublicUser} "User object"
// @Failure 404 {object} domain.ErrorResponse "Not Found - User not found"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Router /user/{identifier} [get]
func (uc *UserController) GetUser(c *gin.Context) {
	identifier := c.Param("identifier")
	user, err := uc.UserUsecase.GetByIdentifier(c, identifier)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "User not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, parser.ToSuccessResponse(user))
}

// @Summary Update user
// @Description Update user by ID
// @Tags User
// @ID updateUser
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body domain.User true "User object"
// @Success 200 "User updated successfully"
// @Failure 400 {object} domain.ErrorResponse "Bad Request - Invalid input"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Router /user/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	identifier := c.Param("id")
	id, err := internal.ParseUint(identifier)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid user ID: " + err.Error(),
		})
		return
	}

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	err = uc.UserUsecase.Update(c, id, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to update user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Delete user
// @Description Delete user by ID
// @Tags User
// @ID deleteUser
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "User deleted successfully"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Router /user/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	identifier := c.Param("id")
	id, err := internal.ParseUint(identifier)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid user ID: " + err.Error(),
		})
		return
	}

	err = uc.UserUsecase.Archive(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to delete user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

package controller

import (
	"golang-chapter-39/LA-Chapter-39H/models"
	"golang-chapter-39/LA-Chapter-39H/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	service *service.UserService
	logger  *zap.Logger
}

func NewUserController(service *service.UserService, logger *zap.Logger) *UserController {
	return &UserController{service, logger}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.CreateUser(user); err != nil {
		ctrl.logger.Error("Failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctrl.logger.Info("User created successfully", zap.Any("user", user))
	c.JSON(http.StatusOK, user)
}

// Similarly, implement other handlers like GetUser, UpdateUser, DeleteUser

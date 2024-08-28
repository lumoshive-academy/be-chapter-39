package controller

import (
	"golang-chapter-39/LA-Chapter-39H-I/models"
	"golang-chapter-39/LA-Chapter-39H-I/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HTTPResponse struct {
	Status      bool        `json:"status"`
	ErrorCode   string      `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func responseOK(c *gin.Context, data interface{}, description string) {
	c.JSON(http.StatusOK, HTTPResponse{
		Status:      true,
		Description: description,
		Data:        data,
	})
}

func responseError(c *gin.Context, errorCode string, description string, httpStatusCode int) {
	c.JSON(httpStatusCode, HTTPResponse{
		Status:      false,
		ErrorCode:   errorCode,
		Description: description,
	})
}

type UserController struct {
	service service.UserService
	logger  *zap.Logger
}

func NewUserController(service service.UserService, logger *zap.Logger) *UserController {
	return &UserController{service, logger}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ctrl.logger.Error("Failed to bind user data", zap.Error(err))
		responseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.CreateUser(user); err != nil {
		ctrl.logger.Error("Failed to create user", zap.Error(err))
		responseError(c, "CREATE_ERROR", "Failed to create user", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("User created successfully", zap.Any("user", user))
	responseOK(c, user, "User created successfully")
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Invalid user ID", zap.Error(err))
		responseError(c, "INVALID_ID", "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := ctrl.service.GetUser(uint(id))
	if err != nil {
		ctrl.logger.Error("User not found", zap.Error(err))
		responseError(c, "NOT_FOUND", "User not found", http.StatusNotFound)
		return
	}

	ctrl.logger.Info("User retrieved successfully", zap.Any("user", user))
	responseOK(c, user, "User retrieved successfully")
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Invalid user ID", zap.Error(err))
		responseError(c, "INVALID_ID", "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ctrl.logger.Error("Failed to bind user data", zap.Error(err))
		responseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = uint(id)

	if err := ctrl.service.UpdateUser(user); err != nil {
		ctrl.logger.Error("Failed to update user", zap.Error(err))
		responseError(c, "UPDATE_ERROR", "Failed to update user", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("User updated successfully", zap.Any("user", user))
	responseOK(c, user, "User updated successfully")
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Invalid user ID", zap.Error(err))
		responseError(c, "INVALID_ID", "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := ctrl.service.DeleteUser(uint(id)); err != nil {
		ctrl.logger.Error("Failed to delete user", zap.Error(err))
		responseError(c, "DELETE_ERROR", "Failed to delete user", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("User deleted successfully", zap.Int("userID", id))
	responseOK(c, nil, "User deleted successfully")
}

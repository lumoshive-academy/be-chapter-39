package controller

import (
	"golang-chapter-39/LA-Chapter-39H/service"

	"go.uber.org/zap"
)

type Controller struct {
	User UserController
}

func NewController(service *service.UserService, logger *zap.Logger) *Controller {
	return &Controller{
		User: *NewUserController(service, logger),
	}
}

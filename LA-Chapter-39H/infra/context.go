package infra

import (
	"fmt"
	"golang-chapter-39/LA-Chapter-39H/config"
	"golang-chapter-39/LA-Chapter-39H/controller"
	"golang-chapter-39/LA-Chapter-39H/database"
	"golang-chapter-39/LA-Chapter-39H/log"
	"golang-chapter-39/LA-Chapter-39H/repository"
	"golang-chapter-39/LA-Chapter-39H/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cfg config.Config
	Ctl controller.Controller
	Log *zap.Logger
}

func NewServiceContext() *ServiceContext {

	config := config.LoadConfig()

	log, err := log.InitZapLogger(config)

	if err != nil {
		fmt.Println("can't init service context %w", err.Error())
	}

	db := database.ConnectDB(config)

	repository := repository.NewUserRepository(db)

	service := service.NewUserService(repository)

	Ctl := controller.NewController(service, log)

	return &ServiceContext{Cfg: config, Ctl: *Ctl, Log: log}
}

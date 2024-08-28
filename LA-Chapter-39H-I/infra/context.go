package infra

import (
	"golang-chapter-39/LA-Chapter-39H-I/config"
	"golang-chapter-39/LA-Chapter-39H-I/controller"
	"golang-chapter-39/LA-Chapter-39H-I/database"
	"golang-chapter-39/LA-Chapter-39H-I/log"
	"golang-chapter-39/LA-Chapter-39H-I/repository"
	"golang-chapter-39/LA-Chapter-39H-I/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cfg config.Config
	Ctl controller.Controller
	Log *zap.Logger
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.LoadConfig()
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := log.InitZapLogger(config)
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	// instance repository
	repository := repository.NewRepository(db)

	// instance service
	service := service.NewService(repository)

	// instance controller
	Ctl := controller.NewController(service, log)

	return &ServiceContext{Cfg: config, Ctl: *Ctl, Log: log}, nil
}

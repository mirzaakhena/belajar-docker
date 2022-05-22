package application

import (
	"belajar-docker/domain_order/controller/restapi"
	"belajar-docker/domain_order/gateway/prod"
	"belajar-docker/domain_order/usecase/runordercreate"
	"belajar-docker/shared/driver"
	"belajar-docker/shared/infrastructure/config"
	"belajar-docker/shared/infrastructure/logger"
	"belajar-docker/shared/infrastructure/server"
	"belajar-docker/shared/infrastructure/util"
)

type order struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c order) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewOrder() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData("order", appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandler(log, ":8080", appData)

		datasource := prod.NewGateway(log, appData, cfg)

		return &order{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:                  log,
				Router:               httpHandler.Router,
				RunOrderCreateInport: runordercreate.NewUsecase(datasource),
			},
		}

	}
}

package application

import (
	"belajar-docker/domain_payment/controller/remote"
	"belajar-docker/domain_payment/controller/restapi"
	"belajar-docker/domain_payment/gateway/prod"
	"belajar-docker/domain_payment/usecase/runpaymentcreate"
	"belajar-docker/shared/driver"
	"belajar-docker/shared/infrastructure/config"
	"belajar-docker/shared/infrastructure/logger"
	"belajar-docker/shared/infrastructure/remoting"
	"belajar-docker/shared/infrastructure/server"
	"belajar-docker/shared/infrastructure/util"
)

type payment struct {
	httpHandler *server.GinHTTPHandler
	Remoting    *remoting.RemoteListener
	controller  driver.Controller
	remoteCall  driver.Controller
}

func (c payment) RunApplication() {
	c.controller.RegisterRouter()
	c.remoteCall.RegisterRouter()
	go c.Remoting.Run()
	c.httpHandler.RunApplication()
}

func NewPayment() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData("payment", appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandler(log, ":8081", appData)

		datasource := prod.NewGateway(log, appData, cfg)

		rem := remoting.NewRemoteListener(1234)

		return &payment{
			httpHandler: &httpHandler,
			Remoting:    rem,
			controller: &restapi.Controller{
				Log:                    log,
				Router:                 httpHandler.Router,
				RunPaymentCreateInport: runpaymentcreate.NewUsecase(datasource),
			},
			remoteCall: &remote.Controller{
				Log:                    log,
				Remote:                 rem,
				RunPaymentCreateInport: runpaymentcreate.NewUsecase(datasource),
			},
		}

	}
}

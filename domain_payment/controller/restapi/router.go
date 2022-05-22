package restapi

import (
	"github.com/gin-gonic/gin"

	"belajar-docker/domain_payment/usecase/runpaymentcreate"
	"belajar-docker/shared/infrastructure/config"
	"belajar-docker/shared/infrastructure/logger"
)

type Controller struct {
	Router                 gin.IRouter
	Config                 *config.Config
	Log                    logger.Logger
	RunPaymentCreateInport runpaymentcreate.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/payment", r.authorized(), r.runPaymentCreateHandler())
}

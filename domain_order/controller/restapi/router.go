package restapi

import (
	"github.com/gin-gonic/gin"

	"belajar-docker/domain_order/usecase/runordercreate"
	"belajar-docker/shared/infrastructure/config"
	"belajar-docker/shared/infrastructure/logger"
)

type Controller struct {
	Router               gin.IRouter
	Config               *config.Config
	Log                  logger.Logger
	RunOrderCreateInport runordercreate.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/order", r.authorized(), r.runOrderCreateHandler(r.RunOrderCreateInport))
}

package remote

import (
	"belajar-docker/domain_payment/usecase/runpaymentcreate"
	"belajar-docker/shared/infrastructure/logger"
	"belajar-docker/shared/infrastructure/remoting"
)

type Controller struct {
	Log                    logger.Logger
	Remote                 *remoting.RemoteListener
	RunPaymentCreateInport runpaymentcreate.Inport
}

func (r *Controller) RegisterRouter() {
	r.Remote.SetHandler(r)
}

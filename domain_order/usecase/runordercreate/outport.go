package runordercreate

import (
	"belajar-docker/domain_order/model/repository"
	"belajar-docker/domain_payment/model/service"
)

// Outport of usecase
type Outport interface {
	repository.SaveOrderRepo
	service.CallRunPaymentCreateService
}

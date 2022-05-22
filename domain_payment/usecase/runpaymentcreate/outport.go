package runpaymentcreate

import "belajar-docker/domain_payment/model/repository"

// Outport of usecase
type Outport interface {
	repository.SavePaymentRepo
}

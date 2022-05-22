package runpaymentcreate

import (
	"belajar-docker/domain_payment/model/entity"
	"context"
)

//go:generate mockery --name Outport -output mocks/

type runPaymentCreateInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runPaymentCreateInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runPaymentCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	err = r.outport.SavePayment(ctx, &entity.Payment{ID: req.ID})
	if err != nil {
		return nil, err
	}

	return res, nil
}

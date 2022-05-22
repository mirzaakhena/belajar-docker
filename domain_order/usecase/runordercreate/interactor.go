package runordercreate

import (
	"belajar-docker/domain_order/model/entity"
	"context"
)

//go:generate mockery --name Outport -output mocks/

type runOrderCreateInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runOrderCreateInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runOrderCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	err = r.outport.SaveOrder(ctx, &entity.Order{ID: req.ID})
	if err != nil {
		return nil, err
	}

	_, err = r.outport.CallRunPaymentCreate(ctx, req.RunPaymentCreateRequest)
	if err != nil {
		return nil, err
	}

	return res, nil
}

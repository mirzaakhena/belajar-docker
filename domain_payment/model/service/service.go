package service

import (
	"belajar-docker/shared/model/remoteparam"
	"context"
)

type CallRunPaymentCreateService interface {
	CallRunPaymentCreate(ctx context.Context, req remoteparam.RunPaymentCreateRequest) (*remoteparam.RunPaymentCreateResponse, error)
}

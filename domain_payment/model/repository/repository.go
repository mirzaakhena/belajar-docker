package repository

import (
	"belajar-docker/domain_payment/model/entity"
	"context"
)

type SavePaymentRepo interface {
	SavePayment(ctx context.Context, obj *entity.Payment) error
}

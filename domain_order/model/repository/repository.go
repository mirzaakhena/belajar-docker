package repository

import (
	"belajar-docker/domain_order/model/entity"
	"context"
)

type SaveOrderRepo interface {
	SaveOrder(ctx context.Context, obj *entity.Order) error
}

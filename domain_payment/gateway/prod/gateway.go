package prod

import (
	"belajar-docker/domain_payment/model/entity"
	"belajar-docker/shared/driver"
	"belajar-docker/shared/infrastructure/config"
	"belajar-docker/shared/infrastructure/logger"
	"context"
)

type gateway struct {
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) SavePayment(ctx context.Context, obj *entity.Payment) error {
	r.log.Info(ctx, "payment with id %s is saved", obj.ID)

	return nil
}

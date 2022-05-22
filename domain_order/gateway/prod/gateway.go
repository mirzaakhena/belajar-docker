package prod

import (
	"belajar-docker/domain_order/model/entity"
	"belajar-docker/shared/driver"
	"belajar-docker/shared/infrastructure/config"
	"belajar-docker/shared/infrastructure/logger"
	"belajar-docker/shared/model/remoteparam"
	"context"
	"net/rpc"
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

func (r *gateway) SaveOrder(ctx context.Context, obj *entity.Order) error {
	r.log.Info(ctx, "Order with orderID %s is Saved", obj.ID)

	return nil
}

func (r *gateway) CallRunPaymentCreate(ctx context.Context, req remoteparam.RunPaymentCreateRequest) (res *remoteparam.RunPaymentCreateResponse, err error) {
	r.log.Info(ctx, "called")

	client, err := rpc.DialHTTP("tcp", "apppayment:1234")
	if err != nil {
		return nil, err
	}

	defer func(client *rpc.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	req.TraceID = logger.GetTraceID(ctx)

	if err := client.Call("Controller.RunPaymentCreateHandler", req, &res); err != nil {
		return nil, err
	}

	return nil, nil
}

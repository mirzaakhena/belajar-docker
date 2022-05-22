package remote

import (
	"belajar-docker/domain_payment/usecase/runpaymentcreate"
	"belajar-docker/shared/infrastructure/logger"
	"belajar-docker/shared/infrastructure/util"
	"belajar-docker/shared/model/remoteparam"
	"context"
)

// RunPaymentCreateHandler ...
func (r *Controller) RunPaymentCreateHandler(rpcReq remoteparam.RunPaymentCreateRequest, rpcRes *remoteparam.RunPaymentCreateResponse) error {

	traceID := rpcReq.TraceID

	ctx := logger.SetTraceID(context.Background(), traceID)

	req := runpaymentcreate.InportRequest{
		RunPaymentCreateRequest: rpcReq,
	}

	r.Log.Info(ctx, util.MustJSON(req))

	res, err := r.RunPaymentCreateInport.Execute(ctx, req)
	if err != nil {
		r.Log.Error(ctx, err.Error())
		return err
	}

	rpcRes = &res.RunPaymentCreateResponse

	r.Log.Info(ctx, util.MustJSON(rpcRes))

	return nil

}

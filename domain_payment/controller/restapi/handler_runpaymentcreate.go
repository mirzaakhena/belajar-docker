package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"belajar-docker/domain_payment/usecase/runpaymentcreate"
	"belajar-docker/shared/infrastructure/logger"
	"belajar-docker/shared/infrastructure/util"
	"belajar-docker/shared/model/payload"
)

// runPaymentCreateHandler ...
func (r *Controller) runPaymentCreateHandler() gin.HandlerFunc {

	type request struct {
		ID string `form:"id,default=13" json:"id"`
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.Bind(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req runpaymentcreate.InportRequest
		req.ID = jsonReq.ID

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunPaymentCreateInport.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}

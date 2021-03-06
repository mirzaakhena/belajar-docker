package payload

import (
	"belajar-docker/shared/model/apperror"
)

type Response struct {
	Success      bool        `json:"success"`
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
	TraceID      string      `json:"traceId"`
}

func NewSuccessResponse(data interface{}, traceID string) interface{} {
	var res Response
	res.Success = true
	res.Data = data
	res.TraceID = traceID
	return res
}

func NewErrorResponse(err error, traceID string) interface{} {
	var res Response
	res.Success = false
	res.TraceID = traceID

	et, ok := err.(apperror.ErrorType)
	if !ok {
		res.ErrorCode = "UNDEFINED"
		res.ErrorMessage = err.Error()
		return res
	}

	res.ErrorCode = et.Code()
	res.ErrorMessage = et.Error()
	return res
}

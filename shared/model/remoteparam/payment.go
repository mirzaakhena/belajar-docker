package remoteparam

type RunPaymentCreateRequest struct {
	TraceID string ``                              //
	ID      string `form:"id,default=5" json:"id"` //
}

type RunPaymentCreateResponse struct {
}

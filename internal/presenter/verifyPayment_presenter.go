package presenter

type VerifyPaymentPresenter struct {
	UserID          uint32  `json:"userId"`
	TransactionID   string  `json:"transactionId"`
	TransactionUUID string  `json:"transactionuuId"`
	TotalAmount     float64 `json:"totalAmount"`
	ProductID       uint    `json:"productId"`
	PaymentMode     string  `json:"paymentMode"`
}

type VerifyPaymentList struct {
	PID         string  `json:"pid"`
	SCD         string  `json:"scd"`
	TotalAmount float64 `json:"totalAmount"`
	Status      string  `json:"status"`
	RefID       string  `json:"refId"`
}

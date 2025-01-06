package models

type Payment struct {
	BaseModel
	UserID          uint    `json:"userId"`
	ProductID       uint    `json:"productID"`
	TransactionID   string  `json:"transactionId"`
	TransactionUUID string  `json:"transactionuuId"`
	TotalAmount     float64 `json:"totalAmount"`
	PaymentMode     string  `json:"paymentMode"`
}

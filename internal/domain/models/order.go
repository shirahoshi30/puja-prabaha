package models

type Order struct {
	BaseModel
	UserID      uint    `json:"userId"`
	GrandTotal  float64 `json:"grandTotal"`
	Address     string  `json:"address"`
	PhoneNumber uint    `json:"phoneNumber"`
	PaymentMode string  `json:"paymentMode"`
	Status      string  `json:"status"`
}

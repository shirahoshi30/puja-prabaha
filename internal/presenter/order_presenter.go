package presenter

type OrderCreateUpdatePresenter struct {
	ID          uint    `json:"Id"`
	UserID      uint    `json:"userId"`
	GrandTotal  float64 `json:"grandTotal"`
	Address     string  `json:"address"`
	PhoneNumber uint    `json:"phoneNumber"`
	PaymentMode string  `json:"paymentMode"`
	Status      string  `json:"status"`
}

type OrderListPresenter struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"userId"`
	Status string `json:"status"`
}

type OrderListReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

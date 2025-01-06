package presenter

type OrderProductCreateListPresenter struct {
	ID        uint `json:"id"`
	UserID    uint `json:"userId"`
	ProductId uint `json:"productId"`
	OrderID   uint `json:"orderId"`
}

type OrderProductListReq struct {
	OrderID  int `json:"categoryId"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

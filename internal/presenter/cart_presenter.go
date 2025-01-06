package presenter

type AddToCart struct {
	ID        uint `json:"Id"`
	VarientID uint `json:"varientId" validate:"required"`
	UserID    uint `json:"userId"`
	Quantity  uint `json:"quantity"`
}

type UpdateCart struct {
	ID        uint `json:"Id"`
	Quantity  uint `json:"quantity"`
	VarientID uint `json:"varientId"`
	UserID    uint `json:"userId"`
}

type CartListPresenter struct {
	ID          uint    `json:"Id"`
	ProductName string  `json:"productName"`
	Quantity    uint    `json:"quantity"`
	TotalPrice  float32 `json:"totalPrice"`
}

type CartDetailPresenter struct {
	Carts      []CartListPresenter `json:"carts"`
	GrandTotal float32             `json:"grandTotal"`
}

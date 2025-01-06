package models

type OrderProduct struct {
	BaseModel
	UserID    uint `json:"userId"`
	OrderID   uint `json:"orderId"`
	ProductID uint `json:"productId"`
}

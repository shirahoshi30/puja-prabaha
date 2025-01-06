package models

type VarientProduct struct {
	BaseModel
	Size            string  `json:"size"`
	Color           string  `json:"color"`
	Price           float32 `json:"price"`
	DiscountedPrice float32 `json:"discount"`
	ProductID       uint    `json:"productId"`
	Product         Product `json:"-"`
}

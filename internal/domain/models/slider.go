package models

type Slider struct {
	BaseModel
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Tag         string  `json:"tag"`
	ProductID   uint    `json:"productId"`
	Product     Product `json:"-"`
}

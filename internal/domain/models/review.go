package models

type Review struct {
	BaseModel
	Rating    float64 `json:"rating"`
	Message   string  `json:"message"`
	ProductID uint    `json:"productId"`
	Product   Product `json:"-"`
	UserID    uint    `json:"userId"`
	User      User    `json:"-"`
}

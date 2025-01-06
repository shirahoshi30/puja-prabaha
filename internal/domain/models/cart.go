package models

type Cart struct {
	BaseModel
	Quantity       uint           `json:"quantity"`
	VarientID      uint           `json:"varientId"`
	VarientProduct VarientProduct `gorm:"foreignKey:VarientID"`
	UserID         uint           `json:"userId"`
}

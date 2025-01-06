package models

import "time"

type BaseModel struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type Product struct {
	BaseModel
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	CategoryID      int              `json:"categoryId"`
	Category        Category         `json:"-"`
	SKU             string           `json:"sku"`
	Tags            string           `json:"tags"`
	VarientProducts []VarientProduct `json:"varientProducts"`
	Reviews         []Review         `json:"-"`
}

package models

type Category struct {
	BaseModel
	Name     string    `json:"name"`
	Products []Product `json:"-"`
}

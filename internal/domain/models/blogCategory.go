package models

type BlogCategory struct {
	BaseModel
	Name string `json:"name"`
	Blog []Blog `json:"-"`
}

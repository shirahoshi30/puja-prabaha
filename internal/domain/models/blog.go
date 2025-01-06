package models

type Blog struct {
	BaseModel
	Title          string `json:"title"`
	Image          string `json:"images"`
	Tag            string `json:"tag"`
	Description    string `json:"description"`
	BlogCategoryID uint   `json:"blogCategoryId"`
	ReadCount      uint   `json:"readCount"`
}

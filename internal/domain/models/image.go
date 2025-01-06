package models

type Image struct {
	BaseModel
	URL    string `json:"url"`
	BlogID uint   `json:"blogId"`
	Blog   Blog   `json:"-"`
}

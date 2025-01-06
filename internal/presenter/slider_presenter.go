package presenter

import "mime/multipart"

type SliderCreateUpdatePresenter struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Image       *multipart.FileHeader `json:"image"`
	Description string                `json:"description"`
	Tag         string                `json:"tag"`
	ProductID   uint                  `json:"productId"`
}

type SliderListPresenter struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	ProductID   uint   `json:"productId"`
	Image       string `json:"image"`
}

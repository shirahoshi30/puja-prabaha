package presenter

import "mime/multipart"

type BlogCreateUpdatePresenter struct {
	ID             uint                    `json:"id"`
	Title          string                  `json:"title"`
	Tag            string                  `json:"tag"`
	Image          []*multipart.FileHeader `json:"images"`
	Description    string                  `json:"description"`
	BlogCategoryID uint                    `json:"blogCategoryId"`
}

type BlogListPresenter struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Tag   string `json:"tag"`
}

type BlogDetailsPresenter struct {
	ID           uint                   `json:"id"`
	Title        string                 `json:"title"`
	Tag          string                 `json:"tag"`
	Description  string                 `json:"description"`
	BlogCategory map[string]interface{} `json:"blogCategory"`
	ReadCount    uint                   `json:"readCount"`
	Image        string                 `json:"image"`
}

type BlogListReq struct {
	BlogCategoryID int `json:"blogCategoryId"`
	Page           int `json:"page"`
	PageSize       int `json:"pageSize"`
}

package presenter

type BlogCategoryPresenter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type BlogCategoryListReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

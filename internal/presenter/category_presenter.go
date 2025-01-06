package presenter

type CategoryPresenter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoryListReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

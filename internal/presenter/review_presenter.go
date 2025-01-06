package presenter

type ReviewList struct {
	Page      int  `json:"page"`
	PageSize  int  `json:"pageSize"`
	ProductID uint `json:"productId"`
}

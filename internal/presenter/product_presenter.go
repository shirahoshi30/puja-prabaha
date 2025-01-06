package presenter

type ProductCreation struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float32 `json:"price"`
	DiscountedPrice float32 `json:"discount"`
	CategoryID      int     `json:"categoryid"`
	SKU             string  `json:"sku"`
	Tags            string  `json:"tags"`
}

type UpdateProduct struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float32 `json:"price"`
	DiscountedPrice float32 `json:"discountedprice"`
	CategoryID      int     `json:"categoryid"`
}

type ProductListPresenter struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Price           float32 `json:"price"`
	DiscountedPrice float32 `json:"discountedPrice"`
	Rating          float32 `json:"rating"`
}

type ProductDetailsPresenter struct {
	ID          uint                        `json:"id"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	SKU         string                      `json:"sku"`
	Tags        string                      `json:"tags"`
	Category    map[string]interface{}      `json:"category"`
	Variants    map[string][]VarientListRes `json:"variants"`
	Rating      float32                     `json:"rating"`
}

type ProductListReq struct {
	CategoryID int `json:"categoryId"`
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
}

type ProductComparePresenter struct {
	Product1 ProductDetailsPresenter `json:"product1"`
	Product2 ProductDetailsPresenter `json:"product2"`
}

type ProductCompareRequest struct {
	ProductIDs []int `json:"productIDs" validate:"min=2"`
}

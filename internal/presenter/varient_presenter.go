package presenter

type VarientProductCreation struct {
	ID              uint    `json:"Id"`
	Size            string  `json:"size"`
	Color           string  `json:"color"`
	Price           float32 `json:"price"`
	DiscountedPrice float32 `json:"discount"`
	ProductID       int     `json:"productId"`
}

type UpdateVarient struct {
	ID              uint    `json:"Id"`
	Size            string  `json:"size"`
	Color           string  `json:"color"`
	Price           float32 `json:"price"`
	DiscountedPrice float32 `json:"discountedprice"`
	ProductID       int     `json:"productId"`
}

type VarientListRes struct {
	ID        uint    `json:"Id"`
	Name      string  `json:"name"`
	ProductID uint    `json:"productId"`
	Color     string  `json:"color"`
	Size      string  `json:"-"`
	Price     float32 `json:"price"`
}

type VariantListReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Max      uint32 `json:"max"`
	Min      uint32 `json:"min"`
	Search   string `json:"search"`
}

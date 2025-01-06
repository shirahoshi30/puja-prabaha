package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) ListVarient(req presenter.VariantListReq) ([]presenter.VarientListRes, int, error) {

	product, totalPage, err := uCase.repo.ListVarient(req)
	if err != nil {
		return nil, int(totalPage), err
	}
	var varientProduct []presenter.VarientListRes
	for i := 0; i < len(product); i++ {

		prodName, err := uCase.repo.GetProductNameWithProductID(product[i].ProductID)
		if err != nil {
			return nil, int(totalPage), err
		}
		varientProduct = append(varientProduct, presenter.VarientListRes{
			ID:        product[i].ID,
			Name:      prodName,
			ProductID: product[i].ProductID,
			Color:     product[i].Color,
			Size:      product[i].Size,
			Price:     product[i].Price,
		})
	}
	// for _, products := range product {
	// 	for _, variant := range products.VarientProducts {
	// 		varientProduct = append(varientProduct, presenter.VarientListRes{
	// 			Price:     variant.Price,
	// 			Color:     variant.Color,
	// 			ProductID: uint(variant.ProductID),
	// 		})
	// 	}
	// }

	// vp, err := json.Marshal(product)
	// if err != nil {
	// 	return nil, err
	// }
	// if err = json.Unmarshal(vp, &varientProduct); err != nil {
	// 	return nil, err
	// }
	return varientProduct, int(totalPage), nil
}

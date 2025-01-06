package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) ListAllProduct(req presenter.ProductListReq) ([]presenter.ProductListPresenter, int, error) {
	products, totalPage, err := uCase.repo.ListAllProduct(req)
	if err != nil {
		return nil, int(totalPage), err
	}
	var allProducts []presenter.ProductListPresenter

	//  variants := make(map[string][]presenter.VarientListRes)

	for _, product := range products {
		eachProduct := presenter.ProductListPresenter{
			ID:   product.ID,
			Name: product.Name,
		}

		if len(product.VarientProducts) != 0 {
			eachProduct.Price = product.VarientProducts[0].Price
			eachProduct.DiscountedPrice = product.VarientProducts[0].DiscountedPrice
			// eachProduct.Rating = float32(product.Reviews[0].Rating)
		}
		// fmt.Printf("product.Reviews: %v\n", product.Reviews)

		if len(product.Reviews) != 0 {
			var ratingSum float32
			for _, review := range product.Reviews {
				ratingSum += float32(review.Rating)
				// ratingSum = ratingSum + float32(review.Rating)
			}
			eachProduct.Rating = ratingSum / float32(len(product.Reviews))
		}

		allProducts = append(allProducts, eachProduct)
	}

	// pp, err := json.Marshal(products)
	// if err != nil {
	// 	return nil, int(totalPage), err
	// }
	// if err = json.Unmarshal(pp, &allProducts); err != nil {
	// 	return nil, int(totalPage), err
	// }
	return allProducts, int(totalPage), nil
}

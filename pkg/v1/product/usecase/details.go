package usecase

import "pujaprabha/internal/presenter"

func (uCase *usecase) DetailsOfProduct(id uint) (*presenter.ProductDetailsPresenter, error) {

	product, err := uCase.repo.DetailsOfProduct(id)
	if err != nil {
		return nil, err
	}

	// var detail presenter.ProductDetailsPresenter
	detail := &presenter.ProductDetailsPresenter{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		SKU:         product.SKU,
		Tags:        product.Tags,
	}

	category, _ := uCase.categoryRepo.GetCategoryByID(uint(product.CategoryID))

	categoryMap := make(map[string]interface{})

	categoryMap["id"] = category.ID
	categoryMap["name"] = category.Name
	detail.Category = categoryMap

	variants := make(map[string][]presenter.VarientListRes)

	for _, variant := range product.VarientProducts {
		var eachVariant presenter.VarientListRes
		var variantList []presenter.VarientListRes
		eachVariant.ID = variant.ID

		eachVariant.ProductID = uint(variant.ProductID)
		eachVariant.Color = variant.Color
		eachVariant.Price = variant.Price
		eachVariant.Name = product.Name
		// eachVariant.Size = variant.Size
		if _, ok := variants[variant.Size]; ok {
			variants[variant.Size] = append(variants[variant.Size], eachVariant)
		} else {
			variantList = append(variantList, eachVariant)
			variants[variant.Size] = variantList
		}

	}
	if len(product.Reviews) != 0 {
		var ratingSum float32
		for _, review := range product.Reviews {
			ratingSum += float32(review.Rating)

		}

		detail.Rating = ratingSum / float32(len(product.Reviews))
	}

	detail.Variants = variants

	return detail, nil
}

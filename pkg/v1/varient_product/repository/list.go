package repository

import (
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
)

func (repo *repository) ListVarient(req presenter.VariantListReq) ([]models.VarientProduct, float64, error) {
	var err error
	// db, _ := config.ConfigDb()
	var product []models.VarientProduct

	totalPage := utils.GetTotalPage(models.VarientProduct{}, repo.db, req.PageSize)

	if totalPage < float64(req.Page) {
		req.Page = int(totalPage)
	}

	if req.Min != 0 && req.Max != 0 {

		err = repo.db.Debug().Model(models.VarientProduct{}).Where("varient_products.price >= ? AND varient_products.price <= ?", req.Min, req.Max).Find(&product).Error
		return product, totalPage, err
	}
	if req.Min != 0 {
		err = repo.db.Debug().Model(models.VarientProduct{}).Where("varient_products.price >= ?", req.Min).Find(&product).Error
		return product, totalPage, err
	}
	if req.Max != 0 {
		err = repo.db.Debug().Model(models.VarientProduct{}).Where("varient_products.price <= ?", req.Max).Find(&product).Error
		return product, totalPage, err
	}

	if req.Search != "" {
		err = repo.db.Debug().Model(models.Product{}).Joins("JOIN varient_products ON products.id = varient_products.product_id").Select("products.*, varient_products.*").
			Where("products.name LIKE ?", "%"+req.Search+"%").Find(&product).Error
		return product, totalPage, err
	}

	err = repo.db.Debug().Model(models.VarientProduct{}).Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&product).Error

	return product, totalPage, err
}

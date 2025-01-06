package repository

import "pujaprabha/internal/domain/models"

func (repo *repository) GetVarientProductByID(i uint) (*models.VarientProduct, error) {

	var varientProduct models.VarientProduct
	err := repo.db.Where("id = ?", i).First(&varientProduct).Error
	if err != nil {

		return nil, err
	}

	return &varientProduct, err
}

func (repo *repository) GetProductNameWithProductID(id uint) (string, error) {
	var product models.Product
	err := repo.db.Select("name").Where("id=?", id).First(&product).Error
	if err != nil {
		return "", err
	}
	return product.Name, err
}

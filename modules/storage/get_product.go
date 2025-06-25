package storage

import (
	"context"
	"myproject/modules/model"
)

func (s *sqlStore) GetProduct(ctx context.Context, cond map[string]interface{}) (*model.ProductWithCatename, error) {
	var data model.ProductWithCatename
	db := s.db.Table("product").
		Select("product.productid, product.productname,product.description, product.price, product.created_at, product.update_at, product.productstatus, categories.categoryname").
		Joins("left join categories on categories.categoryid = product.categoryid").
		Where("productstatus != ?", "Inactive")
	if err := db.Where(cond).First(&data).Error; err != nil {
		return nil, nil
	}
	return &data, nil
}

func (s *sqlStore) GetProductBase(ctx context.Context, cond map[string]interface{}) (*model.Product, error) {
	var data model.Product
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, nil
	}
	return &data, nil
}

package storage

import (
	"context"
	"myproject/common/paging"
	"myproject/modules/model"
)

func (s *sqlStore) GetProducts(ctx context.Context, filter *model.Filter, paging *paging.Paging, moreKeys ...string) ([]model.ProductWithCatename, error) {
	var result []model.ProductWithCatename
	db := s.db.Table("product").
		Select("product.productid, product.productname,product.description, product.price, product.created_at, product.update_at, product.productstatus, categories.categoryname").
		Joins("left join categories on categories.categoryid = product.categoryid").
		Where("productstatus != ?", "Inactive")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("productstatus = ?", v)
		}
	}
	if err := db.Table(model.Product{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("productid desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

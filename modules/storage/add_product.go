package storage

import (
	"context"
	"myproject/modules/model"
)

// Func AddProduct gắn với struct sqlStore có ctx để truyền context từ các tầng trên xuống, data là con trỏ đến struct chứa dữ liệu sản phẩm cần thêm
func (s *sqlStore) AddProduct(ctx context.Context, data *model.AddProduct) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

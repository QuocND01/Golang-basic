package storage

import (
	"context"
	"myproject/modules/model"
)

func (s *sqlStore) UpdateProduct(ctx context.Context, cond map[string]interface{}, updata *model.UpdateProduct) error {
	if err := s.db.Where(cond).Updates(updata).Error; err != nil {
		return err
	}
	return nil
}

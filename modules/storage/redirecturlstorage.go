package storage

import (
	"context"
	"myproject/modules/model"
)

func (s *sqlStore) GetOriginURL(ctx context.Context, code string) (string, error) {
	var url model.Url
	if err := s.db.WithContext(ctx).
		Where("sorturl = ?", code).
		First(&url).Error; err != nil {
		return "", err
	}
	return url.Originurl, nil
}

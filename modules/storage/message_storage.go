package storage

import (
	"context"
	"myproject/modules/model"
)

func (s *sqlStore) SaveMessage(ctx context.Context, msg *model.Message) error {
	return s.db.WithContext(ctx).Create(msg).Error
}

func (s *sqlStore) ListMessages(ctx context.Context) ([]model.Message, error) {
	var msgs []model.Message
	err := s.db.WithContext(ctx).Find(&msgs).Error
	return msgs, err
}

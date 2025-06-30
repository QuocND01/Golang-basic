package storage

import (
	"context"
	"math/rand"
	"myproject/modules/model"
	"time"
)

func (s *sqlStore) SortenUrl(ctx context.Context, url string, length int) (string, error) {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b), nil
}

func (s *sqlStore) Createurl(ctx context.Context, url *model.Urladd) error {
	if err := s.db.Create(url).Error; err != nil {
		return err
	}
	return nil
}

package biz

import (
	"context"
	"encoding/json"
	"myproject/consumer/model"
)

type ArticleStorage interface {
	Save(article *model.Article) error
}

type KafkaReader interface {
	ReadMessage(ctx context.Context) ([]byte, error)
}

type ConsumeBiz struct {
	store ArticleStorage
}

func NewConsumeBiz(store ArticleStorage) *ConsumeBiz {
	return &ConsumeBiz{store: store}
}

func (biz *ConsumeBiz) ConsumeLoop(ctx context.Context, reader KafkaReader) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			data, err := reader.ReadMessage(ctx)
			if err != nil {
				return err
			}
			var article model.Article
			if err := json.Unmarshal(data, &article); err == nil {
				_ = biz.store.Save(&article)
			}
		}
	}
}

package storage

import (
	"context"
	"encoding/json"
	"myproject/modules/model"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(broker, topic string) *KafkaProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &KafkaProducer{writer: w}
}

func (p *KafkaProducer) PublishArticle(ctx context.Context, article *model.Article) error {
	data, err := json.Marshal(article)
	if err != nil {
		return err
	}
	return p.writer.WriteMessages(ctx, kafka.Message{Value: data})
}

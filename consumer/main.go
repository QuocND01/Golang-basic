package main

import (
	"context"
	"log"
	"myproject/consumer/biz"
	"myproject/consumer/storage"
)

func main() {
	dsn := "newsuser:newspass@tcp(mysql:3306)/newsdb?parseTime=true&charset=utf8mb4"
	db := storage.NewMySQLDB(dsn)
	repo := storage.NewArticleMySQLStorage(db)
	biz := biz.NewConsumeBiz(repo)

	consumer := storage.NewKafkaConsumer("kafka:9092", "news.articles")

	log.Println("Consumer started...")
	if err := biz.ConsumeLoop(context.Background(), consumer); err != nil {
		log.Fatal("consume error:", err)
	}
}

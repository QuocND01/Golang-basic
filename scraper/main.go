package main

import (
	"context"
	"log"
	"myproject/scraper/biz"
	"myproject/scraper/storage"
	"myproject/scraper/transport"
)

func main() {
	producer := storage.NewKafkaProducer("kafka:9092", "news.articles")
	biz := biz.NewScrapeBiz(producer)

	sources := []string{
		"https://vnexpress.net/rss/tin-moi-nhat.rss",
		"https://rss.cnn.com/rss/edition.rss",
		"https://www.theguardian.com/world/rss",
	}

	ctx := context.Background()
	err := transport.StartScraping(ctx, biz, sources)
	if err != nil {
		log.Fatal("scraping error:", err)
	}
}

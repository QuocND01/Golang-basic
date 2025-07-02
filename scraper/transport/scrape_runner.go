package transport

import (
	"context"
	"log"
	"myproject/scraper/biz"
	"sync"
	"time"
)

func StartScraping(ctx context.Context, biz *biz.ScrapeBiz, sources []string) error {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	scrapeAll := func() {
		var wg sync.WaitGroup
		for _, src := range sources {
			src := src
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := biz.ScrapeFeed(ctx, src); err != nil {
					log.Println("scrape error:", err)
				}
			}()
		}
		wg.Wait()
	}

	scrapeAll()
	for {
		select {
		case <-ticker.C:
			scrapeAll()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

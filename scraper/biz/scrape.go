package biz

import (
	"context"
	"myproject/modules/model"

	"github.com/mmcdole/gofeed"
)

type ScraperStorage interface {
	PublishArticle(ctx context.Context, data *model.Article) error
}

type ScrapeBiz struct {
	store ScraperStorage
}

func NewScrapeBiz(store ScraperStorage) *ScrapeBiz {
	return &ScrapeBiz{store: store}
}

func (biz *ScrapeBiz) ScrapeFeed(ctx context.Context, feedURL string) error {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(feedURL)
	if err != nil {
		return err
	}
	for _, item := range feed.Items {
		article := &model.Article{
			Title:     item.Title,
			Content:   item.Description,
			URL:       item.Link,
			Published: item.Published,
			Source:    feedURL,
		}
		_ = biz.store.PublishArticle(ctx, article)
	}
	return nil
}

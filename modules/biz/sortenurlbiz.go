package biz

import (
	"context"
	"myproject/modules/model"
)

type UrlShortenStogare interface {
	SortenUrl(ctx context.Context, url string, length int) (string, error)
	Createurl(ctx context.Context, url *model.Urladd) error
}

type urlShortenBiz struct {
	store UrlShortenStogare
}

func UrlShortenBiz(store UrlShortenStogare) *urlShortenBiz {
	return &urlShortenBiz{store: store}
}

func (biz *urlShortenBiz) UrlShortenandCreate(ctx context.Context, url string, length int) (string, error) {
	code, err := biz.store.SortenUrl(ctx, url, length)
	if err != nil {
		return "", err
	}
	newUrl := &model.Urladd{
		Originurl: url,
		Sorturl:   code,
	}
	if err := biz.store.Createurl(ctx, newUrl); err != nil {
		return "", err
	}
	return code, nil
}

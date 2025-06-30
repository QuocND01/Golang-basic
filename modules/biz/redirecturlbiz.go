package biz

import (
	"context"
)

type Redirecturlstogare interface {
	GetOriginURL(ctx context.Context, code string) (string, error)
}

type redirecturlbiz struct {
	store Redirecturlstogare
}

func Redirecturlbiz(store Redirecturlstogare) *redirecturlbiz {
	return &redirecturlbiz{store: store}
}

func (biz *redirecturlbiz) GetOriginURL(ctx context.Context, code string) (string, error) {
	origin, err := biz.store.GetOriginURL(ctx, code)
	if err != nil {
		return "", err
	}
	return origin, nil
}

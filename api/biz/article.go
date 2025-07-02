package biz

import "myproject/api/model"

type ArticleStorage interface {
	ListArticles() ([]*model.Article, error)
}

type ArticleBiz struct {
	store ArticleStorage
}

func NewArticleBiz(store ArticleStorage) *ArticleBiz {
	return &ArticleBiz{store: store}
}

func (b *ArticleBiz) GetLatestArticles() ([]*model.Article, error) {
	return b.store.ListArticles()
}

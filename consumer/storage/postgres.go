package storage

import (
	"database/sql"
	"log"
	"myproject/consumer/model"
)

type ArticlePostgresStorage struct {
	db *sql.DB
}

func NewPostgresDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Postgres connect error:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Postgres ping error:", err)
	}
	return db
}

func NewArticlePostgresStorage(db *sql.DB) *ArticlePostgresStorage {
	return &ArticlePostgresStorage{db: db}
}

func (s *ArticlePostgresStorage) Save(article *model.Article) error {
	_, err := s.db.Exec(`
		INSERT INTO articles (title, content, url, published, source)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (url) DO NOTHING
	`, article.Title, article.Content, article.URL, article.Published, article.Source)
	return err
}

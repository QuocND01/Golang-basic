package storage

import (
	"database/sql"
	"log"
	"myproject/consumer/model"

	_ "github.com/go-sql-driver/mysql"
)

type ArticleMySQLStorage struct {
	db *sql.DB
}

func NewMySQLDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("MySQL connect error:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("MySQL ping error:", err)
	}
	return db
}

func NewArticleMySQLStorage(db *sql.DB) *ArticleMySQLStorage {
	return &ArticleMySQLStorage{db: db}
}

func (s *ArticleMySQLStorage) Save(article *model.Article) error {
	_, err := s.db.Exec(`
		INSERT INTO articles (title, content, url, published, source)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE title=VALUES(title)
	`, article.Title, article.Content, article.URL, article.Published, article.Source)
	return err
}

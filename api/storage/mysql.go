package storage

import (
	"database/sql"
	"log"
	"myproject/api/model"

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

func (s *ArticleMySQLStorage) ListArticles() ([]*model.Article, error) {
	rows, err := s.db.Query("SELECT title, content, url, published, source FROM articles ORDER BY published DESC LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*model.Article
	for rows.Next() {
		a := new(model.Article)
		if err := rows.Scan(&a.Title, &a.Content, &a.URL, &a.Published, &a.Source); err != nil {
			continue
		}
		articles = append(articles, a)
	}
	return articles, nil
}

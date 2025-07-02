package main

import (
	"log"
	"myproject/api/biz"
	"myproject/api/storage"
	"myproject/api/transport"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	dsn := "newsuser:newspass@tcp(mysql:3306)/newsdb?parseTime=true&charset=utf8mb4"
	db := storage.NewMySQLDB(dsn)

	repo := storage.NewArticleMySQLStorage(db)
	biz := biz.NewArticleBiz(repo)
	handler := transport.NewAPIHandler(biz, rdb)

	r := gin.Default()
	r.GET("/articles", handler.ListArticlesGin)

	log.Println("API server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}

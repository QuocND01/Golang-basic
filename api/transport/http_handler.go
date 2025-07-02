package transport

import (
	"context"
	"myproject/api/biz"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type APIHandler struct {
	biz *biz.ArticleBiz
	rdb *redis.Client
}

func NewAPIHandler(biz *biz.ArticleBiz, rdb *redis.Client) *APIHandler {
	return &APIHandler{biz: biz, rdb: rdb}
}

func (h *APIHandler) ListArticlesGin(c *gin.Context) {
	ctx := context.Background()
	ip := c.ClientIP()
	key := "rate:" + ip

	count, _ := h.rdb.Incr(ctx, key).Result()
	if count == 1 {
		h.rdb.Expire(ctx, key, 10*time.Second)
	}
	if count > 5 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
		return
	}

	articles, err := h.biz.GetLatestArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

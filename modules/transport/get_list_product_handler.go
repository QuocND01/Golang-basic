package transport

import (
	"encoding/json"
	"fmt"
	"myproject/common/logger"
	"myproject/common/paging"
	"myproject/common/redis"
	"myproject/common/response"
	"myproject/modules/biz"
	"myproject/modules/model"
	"myproject/modules/storage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetListProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging paging.Paging
		if err := c.ShouldBind(&paging); err != nil {
			logger.Logger.Error("Failed to bind GetListProduct request", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		paging.Process()

		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			logger.Logger.Error("Failed to bind filter request", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		cacheKey := fmt.Sprintf("product:list:page=%d:limit=%d:filter=%v", paging.Page, paging.Limit, filter)

		// 1. Kiểm tra Redis
		val, err := redis.Client.Get(c.Request.Context(), cacheKey).Result()
		if err == nil {
			var cachedData []model.Product
			if jsonErr := json.Unmarshal([]byte(val), &cachedData); jsonErr == nil {
				logger.Logger.Info("Cache hit - returning data from Redis")
				c.JSON(http.StatusOK, response.NewSuccessResponese(cachedData, paging, filter))
				return
			}
		}

		store := storage.NewSQLStore(db)
		busines := biz.GetListProductBiz(store)
		data, err := busines.GetProducts(c.Request.Context(), &filter, &paging)
		if err != nil {
			logger.Logger.Error("Failed to get list of products", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		// 2. Cache kết quả
		jsonData, _ := json.Marshal(data)
		redis.Client.Set(c.Request.Context(), cacheKey, jsonData, time.Minute*5)

		logger.Logger.Info("Get list of products successfully", zap.Any("data", data))
		c.JSON(http.StatusOK, response.NewSuccessResponese(data, paging, filter))
	}
}

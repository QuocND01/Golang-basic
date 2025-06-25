package transport

import (
	"encoding/json"
	"fmt"
	"myproject/common/logger"
	"myproject/common/redis"
	"myproject/common/response"
	"myproject/modules/biz"
	"myproject/modules/model"
	"myproject/modules/storage"
	"time"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetProductByID(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("productid"))
		if err != nil {
			logger.Logger.Error("Failed to atoi product by productid", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		cacheKey := fmt.Sprintf("product:%d", id)
		val, err := redis.Client.Get(c.Request.Context(), cacheKey).Result()
		if err == nil {
			var cachedData model.Product
			if jsonErr := json.Unmarshal([]byte(val), &cachedData); jsonErr == nil {
				logger.Logger.Info("Cache hit - returning product from Redis")
				c.JSON(http.StatusOK, response.SimpleSuccessResponse(cachedData))
				return
			}
		}

		store := storage.NewSQLStore(db)
		business := biz.GetProductBiz(store)

		data, err := business.GetProductByID(c.Request.Context(), id)
		if err != nil {
			logger.Logger.Error("Failed to get product by productid ", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		jsonData, _ := json.Marshal(data)
		redis.Client.Set(c.Request.Context(), cacheKey, jsonData, time.Minute*5)
		logger.Logger.Info("Get product by productid successfully", zap.Any("data", data))
		c.JSON(http.StatusOK, response.SimpleSuccessResponse(data))
	}
}

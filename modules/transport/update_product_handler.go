package transport

import (
	"fmt"
	"myproject/common/logger"
	"myproject/common/redis"
	"myproject/common/response"
	"myproject/modules/biz"
	"myproject/modules/model"
	"myproject/modules/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func UpdateProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UpdateProduct
		id, err := strconv.Atoi(c.Param("productid"))
		if err != nil {
			logger.Logger.Error("Failed to atoi product by productid", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			logger.Logger.Error("Failed to bind UpdateProduct request", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		store := storage.NewSQLStore(db)
		business := biz.UpdateProductBiz(store)
		if err := business.UpdateProduct(c.Request.Context(), id, &data); err != nil {
			logger.Logger.Error("Failed to update product", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		cacheKey := fmt.Sprintf("product:%d", id)
		if err := redis.Client.Del(c.Request.Context(), cacheKey).Err(); err == nil {
			logger.Logger.Info("Cleared Redis cache after update", zap.String("key", cacheKey))
		}
		logger.Logger.Info("Update product successfully")
		c.JSON(http.StatusOK, response.SimpleSuccessResponse(true))
	}
}

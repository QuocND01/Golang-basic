package transport

import (
	"myproject/common/logger"
	"myproject/common/response"
	"myproject/modules/biz"
	"myproject/modules/model"
	"myproject/modules/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func AddProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.AddProduct
		if err := c.ShouldBind(&data); err != nil {
			logger.Logger.Error("Failed to bind AddProduct request", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.AddProductBiz(store)
		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			logger.Logger.Error("Failed to add new product", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		logger.Logger.Info("Successfully created new product", zap.Any("data", data))
		c.JSON(http.StatusOK, response.SimpleSuccessResponse(data))
	}
}

package transport

import (
	"myproject/common/logger"
	"myproject/common/response"
	"myproject/modules/biz"
	"myproject/modules/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateSortenurl(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var req struct {
			OriginUrl string `json:"origin_url" binding:"required,url"`
			Length    int    `json:"length"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if req.Length == 0 {
			req.Length = 6
		}

		store := storage.NewSQLStore(db)
		business := biz.UrlShortenBiz(store)

		code, err := business.UrlShortenandCreate(c.Request.Context(), req.OriginUrl, req.Length)
		if err != nil {
			logger.Logger.Error("Failed to shorten URL", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		shortUrl := "https://yourdomain.com/" + code

		c.JSON(http.StatusOK, response.SimpleSuccessResponse(gin.H{
			"origin_url": req.OriginUrl,
			"short_url":  shortUrl,
		}))
	}
}

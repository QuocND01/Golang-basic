package transport

import (
	"myproject/modules/biz"
	"myproject/modules/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ResolveShortURL(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		store := storage.NewSQLStore(db)
		business := biz.Redirecturlbiz(store)

		origin, err := business.GetOriginURL(c.Request.Context(), code)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.Redirect(http.StatusFound, origin)
	}
}

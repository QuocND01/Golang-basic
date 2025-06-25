package main

import (
	"fmt"
	"log"
	"myproject/common/auth"
	"myproject/common/hub"
	"myproject/common/logger"
	"myproject/common/middleware"
	"myproject/common/redis"
	"myproject/modules/biz"
	"myproject/modules/storage"
	"myproject/modules/transport"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	logger.InitLogger()
	defer logger.Logger.Sync()

	redis.InitRedisClient()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Logger.Fatal("Cannot connect DB", zap.Error(err))
	}
	middleware.InitMetrics()
	r := gin.Default()
	r.Use(middleware.PrometheusMiddleware())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// OAuth2 login
	r.GET("/auth/google/login", auth.GoogleLogin)
	r.GET("/auth/google/callback", auth.GoogleCallback)

	// JWT protected routes
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.GET("/profile", func(c *gin.Context) {
		email, _ := c.Get("email")
		c.JSON(http.StatusOK, gin.H{"email": email})
	})

	{
		v1 := r.Group("/v1")
		{
			Products := v1.Group("/product")
			{
				Products.POST("", transport.AddProduct(db))
				Products.GET("", transport.GetListProduct(db))
				Products.GET("/:productid", transport.GetProductByID(db))
				Products.PATCH("/:productid", transport.UpdateProduct(db))
				// books.DELETE("/:id", transport.DeleteBook(db))
			}

		}
	}

	store := storage.NewSQLStore(db)
	bizLayer := biz.NewMessageBiz(store)
	hub := hub.NewHub()
	go hub.Run()
	handler := transport.NewWSHandler(bizLayer, hub)
	r.GET("/ws", handler.HandleWS)

	logger.Logger.Info("Server started at :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Logger.Fatal("Server error", zap.Error(err))
	}
}

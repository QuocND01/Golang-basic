package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var HttpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"method", "path", "status"},
)

func InitMetrics() {
	prometheus.MustRegister(HttpRequests)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status := fmt.Sprintf("%d", c.Writer.Status())
		HttpRequests.WithLabelValues(c.Request.Method, c.FullPath(), status).Inc()
	}
}

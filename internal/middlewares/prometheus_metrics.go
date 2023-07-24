package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gin_http_request_duration_seconds",
			Help:    "Histogram of HTTP request duration.",
			Buckets: []float64{0.1, 0.3, 1, 3, 5, 10},
		},
		[]string{"service", "route", "method", "status"},
	)

	httpRequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gin_http_request_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"service", "route", "method", "status"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(httpRequestCounter)
}

// PrometheusMetrics is a middleware that instruments the router with Prometheus metrics.
func PrometheusMetrics(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Continue handling the request.
		c.Next()

		duration := time.Since(start)
		route := c.FullPath()
		method := c.Request.Method
		status := strconv.Itoa(c.Writer.Status())

		// Record the duration and counter metrics.
		httpRequestDuration.WithLabelValues(serviceName, route, method, status).Observe(duration.Seconds())
		httpRequestCounter.WithLabelValues(serviceName, route, method, status).Inc()
	}
}

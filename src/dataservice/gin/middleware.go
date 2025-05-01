package gin

import (
	"math"
	"os"

	"golang-backend-microservice/container/log"
	Time "golang-backend-microservice/container/time"

	"github.com/gin-gonic/gin"
)

type GinMiddleware struct {
	time Time.Time
}

func (t GinMiddleware) Logger() gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := t.time.Now()
		c.Next()

		stop := t.time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		log.Info(
			"hostname: %v, status: %v, latency: %v, clientIP: %v, method: %v, path: %v, referer: %v, dataLength: %v, clientUserAgent: %v",
			hostname,
			statusCode,
			latency,
			clientIP,
			c.Request.Method,
			path,
			referer,
			dataLength,
			clientUserAgent,
		)
	}
}

package gin

import (
	"golang-backend-microservice/container/env"
	"golang-backend-microservice/container/time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func SetupRoutes(nc *nats.Conn) *gin.Engine {
	r := gin.New()
	r.SetTrustedProxies(nil)
	gin.SetMode(gin.ReleaseMode)
	if !env.IsEnv(env.ENV_PRODUCTION) {
		gin.SetMode(gin.DebugMode)
		r.SetTrustedProxies([]string{"127.0.0.1"})
	}
	r.Use(GinMiddleware{
		time: time.RealTime{},
	}.Logger())
	r.Use(gin.Recovery())
	r.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	return r
}

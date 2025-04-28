package nats

import (
	"fmt"
	"golang-backend-microservice/container/env"
	"golang-backend-microservice/container/log"
	"net/http"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/rollbar/rollbar-go"
)

type ServiceConfig struct {
	ServiceName  string
	Version      string
	Description  string
	EndpointName string
}

type Connection struct {
	User string
	Pass string
	Host string
	ServiceConfig
}

func OpenNatsServerConnection() (*nats.Conn, micro.Service) {
	// Get package version
	version := os.Getenv("VERSION")
	if !env.IsEnv(env.ENV_PRODUCTION) {
		version += "-" + os.Getenv("ENVIRONMENT")
	}

	// Establish NATS server connection
	return Connection{
		Host: os.Getenv("NATS_HOST"),
		User: os.Getenv("NATS_USER"),
		Pass: os.Getenv("NATS_PASS"),
		ServiceConfig: ServiceConfig{
			ServiceName:  "Database-Backend",
			Version:      version,
			Description:  "Microservice for database requests and responses",
			EndpointName: "db-backend",
		},
	}.Open()
}

func (c Connection) Open() (*nats.Conn, micro.Service) {
	// Add connection
	nc, err := nats.Connect(
		c.Host, nats.UserInfo(c.User, c.Pass),
		nats.PingInterval(20*time.Second), nats.MaxPingsOutstanding(5),
	)
	if err != nil {
		log.Error("Error connecting to NATS - %s", err)
		return nil, nil
	}

	// Add services
	svc, err := micro.AddService(nc, micro.Config{
		Name:        c.ServiceName,
		Version:     c.Version,
		Description: c.Description,
		Endpoint: &micro.EndpointConfig{
			Subject: c.EndpointName,
			Handler: micro.HandlerFunc(func(req micro.Request) {
				req.Respond(fmt.Append(nil, http.StatusOK))
			}),
		},
		ErrorHandler: func(s micro.Service, e *micro.NATSError) {
			log.Error(e.Error())
		},
	})
	if err != nil {
		log.Error("Error adding NATS service - %s", err)
		return nc, nil
	}

	// Log rollbar
	rollbar.SetCodeVersion(c.Version)
	rollbar.SetCustom(map[string]any{
		"ServiceName": c.ServiceName,
		"ServiceID":   svc.Info().ID,
	})

	return nc, svc
}

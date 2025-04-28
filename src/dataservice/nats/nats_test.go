package nats_test

import (
	"fmt"
	"golang-backend-microservice/dataservice/nats"
	"testing"

	"github.com/nats-io/nats-server/v2/server"
	mockserver "github.com/nats-io/nats-server/v2/test"
)

var (
	USER string = "local"
	PASS string = "password"
	PORT int    = 8369
)

func mockNatsServer(user string, pass string, port int) *server.Server {
	options := mockserver.DefaultTestOptions
	options.Port = port
	options.Username = user
	options.Password = pass
	return mockserver.RunServer(&options)
}

func mockNatsConnection(user string, pass string, port int) nats.Connection {
	return nats.Connection{
		User: user,
		Pass: pass,
		Host: fmt.Sprintf("nats://127.0.0.1:%d", port),
		ServiceConfig: nats.ServiceConfig{
			ServiceName:  "Database",
			Version:      "1.0.0",
			Description:  "Lorem ipsum",
			EndpointName: "database",
		},
	}
}

func TestNatsConnection(t *testing.T) {
	server := mockNatsServer(USER, PASS, PORT)
	defer server.Shutdown()

	// Test open NATS connection
	nc, svc := mockNatsConnection(USER, PASS, PORT).Open()
	defer nc.Close()
	if nc == nil || svc == nil {
		t.Errorf("Error: Nats connection unsuccessful")
	}
}

func BenchmarkOpenNatsConnection(b *testing.B) {
	server := mockNatsServer(USER, PASS, PORT)
	defer server.Shutdown()

	for b.Loop() {
		nc, svc := mockNatsConnection(USER, PASS, PORT).Open()
		defer func() {
			nc.Close()
			svc.Stop()
		}()
	}
}

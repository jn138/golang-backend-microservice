package nats

import (
	"encoding/json"
	"golang-backend-microservice/model"
	"time"

	"github.com/nats-io/nats.go"
)

type requestable interface {
	model.MySqlReqArgs
}

func Request[T requestable](nc *nats.Conn, subject string, args T) (*nats.Msg, error) {
	config, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	return nc.Request(subject, config, 1000*time.Millisecond)
}

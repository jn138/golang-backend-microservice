package nats

import (
	"encoding/json"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/model"

	"github.com/nats-io/nats.go/micro"
)

type responsive interface {
	model.Book
}

type StatusResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

type DataResponse[T responsive] struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   []T    `json:"data"`
}

func (res StatusResponse) Respond(req micro.Request) error {
	b, _ := json.Marshal(res)
	if err := req.Respond(b); err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (res DataResponse[T]) Respond(req micro.Request) error {
	b, _ := json.Marshal(res)
	if err := req.Respond(b); err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

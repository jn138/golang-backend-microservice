package main

import (
	"golang-backend-microservice/config"
	"golang-backend-microservice/dataservice"
)

func main() {
	config.Init()
	dataservice.EstablishDataService()
}

package config

import (
	"golang-backend-microservice/container/env"
)

func Init() {
	env.LoadVariables()
}

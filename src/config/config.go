package config

import (
	"golang-backend-microservice/container/env"
	"golang-backend-microservice/container/log"
)

func Init() {
	// Load env variables & initialise logging system
	env.LoadVariables()
	if env.IsEnv(env.ENV_DEVELOPMENT) {
		log.CreateTransports(log.Console)
	} else {
		log.CreateTransports(log.Console, log.File, log.Rollbar)
	}
	log.Info("Configurations initialized successfully!")
}

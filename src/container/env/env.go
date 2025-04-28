package env

import (
	"flag"
	"os"
	"slices"

	"github.com/joho/godotenv"
)

type Env = string

const (
	ENV_DEVELOPMENT Env = "development"
	ENV_TESTING     Env = "testing"
	ENV_STAGING     Env = "staging"
	ENV_PRODUCTION  Env = "production"
)

const (
	// Base
	ENV_CONFIG_FILENAME    = ".env." + ENV_DEVELOPMENT
	ENV_CONFIG_VERSION     = "0.1.0"
	ENV_CONFIG_SERVER_ENV  = ENV_DEVELOPMENT
	ENV_CONFIG_SERVER_PORT = "8080"

	// NATS server
	ENV_CONFIG_NATS_SERVER_HOST    = "localhost:4222"
	ENV_CONFIG_NATS_SERVER_USER    = "local"
	ENV_CONFIG_NATS_SERVER_PASS    = "password"
	ENV_CONFIG_NATS_SERVER_TIMEOUT = "1000"

	// MySQL database
	ENV_CONFIG_MYSQL_DB_HOST = "db"
	ENV_CONFIG_MYSQL_DB_USER = "root"
	ENV_CONFIG_MYSQL_DB_PASS = "password"

	// Others
	ENV_CONFIG_ROLLBAR_ACCESS_TOKEN = ""
)

func IsEnv(envs ...Env) bool {
	return slices.Contains(envs, os.Getenv("ENVIRONMENT"))
}

func LoadVariables() {
	envList := flag.String("env", ".env.development", "Env file")
	if err := godotenv.Load(*envList); err != nil {
		// Return if environment file exists
		_, exists := os.LookupEnv("ENVIRONMENT")
		if exists {
			return
		}

		// Otherwise, create a sample .env.delopment file
		newEnvVariables := map[string]string{
			"VERSION":     ENV_CONFIG_VERSION,
			"ENVIRONMENT": ENV_CONFIG_SERVER_ENV,
			"PORT":        ENV_CONFIG_SERVER_PORT,

			"NATS_HOST":    ENV_CONFIG_NATS_SERVER_HOST,
			"NATS_USER":    ENV_CONFIG_NATS_SERVER_USER,
			"NATS_PASS":    ENV_CONFIG_NATS_SERVER_PASS,
			"NATS_TIMEOUT": ENV_CONFIG_NATS_SERVER_TIMEOUT,

			"MYSQL_HOST": ENV_CONFIG_MYSQL_DB_HOST,
			"MYSQL_USER": ENV_CONFIG_MYSQL_DB_USER,
			"MYSQL_PASS": ENV_CONFIG_MYSQL_DB_PASS,

			"ROLLBAR_ACCESS_TOKEN": ENV_CONFIG_ROLLBAR_ACCESS_TOKEN,
		}
		godotenv.Write(newEnvVariables, ENV_CONFIG_FILENAME)
		godotenv.Load(ENV_CONFIG_FILENAME)
	}
}

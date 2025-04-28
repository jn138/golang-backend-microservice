package env_test

import (
	"golang-backend-microservice/container/env"
	"testing"
)

func TestIsEnv(t *testing.T) {
	t.Setenv("ENVIRONMENT", env.ENV_DEVELOPMENT)
	if !env.IsEnv(env.ENV_DEVELOPMENT) {
		t.Errorf("Expected IsEnv(%s) to be true, but got false", env.ENV_DEVELOPMENT)
	}
	if env.IsEnv(env.ENV_PRODUCTION) {
		t.Errorf("Expected IsEnv(%s) to be false, but got true", env.ENV_PRODUCTION)
	}
}

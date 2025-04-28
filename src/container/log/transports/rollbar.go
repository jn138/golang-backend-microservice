package transport

import (
	"fmt"
	"golang-backend-microservice/container/env"
	"os"

	"github.com/rollbar/rollbar-go"
)

type Rollbar struct {
	Loglevel int32
	Alert    bool
}

const (
	errors  = 1
	warning = 2
	info    = 3
	debug   = 4
)

func (r Rollbar) Default() Rollbar {
	token, exists := os.LookupEnv("ROLLBAR_ACCESS_TOKEN")
	if !exists || token == "" {
		r.Alert = false
	} else {
		r.Alert = true
		rollbar.SetToken(token)
	}

	rollbar.SetEnvironment(os.Getenv("ENVIRONMENT"))
	r.Loglevel = warning
	if env.IsEnv(env.ENV_DEVELOPMENT, env.ENV_TESTING, env.ENV_STAGING) {
		r.Loglevel = debug
	}

	return r
}

func (r Rollbar) Debug(format string, args ...any) {
	if r.Loglevel == debug && r.Alert {
		rollbar.Debug(fmt.Sprintf(format, args...))
	}
}

func (r Rollbar) Info(format string, args ...any) {
	if r.Loglevel >= info && r.Alert {
		rollbar.Info(fmt.Sprintf(format, args...))
	}
}

func (r Rollbar) Warn(format string, args ...any) {
	if r.Loglevel >= warning && r.Alert {
		rollbar.Warning(fmt.Sprintf(format, args...))
	}
}

func (r Rollbar) Error(format string, args ...any) {
	if r.Loglevel >= errors && r.Alert {
		rollbar.Error(fmt.Sprintf(format, args...))
	}
}

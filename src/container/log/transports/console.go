package transport

import (
	"golang-backend-microservice/container/env"
	"golang-backend-microservice/container/time"
	"os"

	log "github.com/sirupsen/logrus"
)

type Console struct {
	Context *log.Entry
}

func (c Console) Default() Console {
	formatter := log.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: time.DATE_TIME_LAYOUT,
	}
	loglevel := log.ErrorLevel

	if env.IsEnv(env.ENV_DEVELOPMENT, env.ENV_TESTING, env.ENV_STAGING) {
		formatter = log.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: time.TIME_LAYOUT_IN_MS,
		}
		loglevel = log.DebugLevel
	}

	c.Context = log.NewEntry(&log.Logger{
		Out:       os.Stderr,
		Formatter: &formatter,
		Level:     loglevel,
		Hooks:     make(log.LevelHooks),
	})

	return c
}

func (c Console) Debug(format string, args ...any) {
	c.Context.Debugf(format, args...)
}

func (c Console) Info(format string, args ...any) {
	c.Context.Infof(format, args...)
}

func (c Console) Warn(format string, args ...any) {
	c.Context.Warningf(format, args...)
}

func (c Console) Error(format string, args ...any) {
	c.Context.Errorf(format, args...)
}

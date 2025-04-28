package transport

import (
	"golang-backend-microservice/container/env"
	"os"

	log "github.com/sirupsen/logrus"
)

type File struct {
	Context *log.Entry
}

func (f File) Default() File {
	loglevel := log.ErrorLevel

	if env.IsEnv(env.ENV_DEVELOPMENT, env.ENV_TESTING, env.ENV_STAGING) {
		loglevel = log.DebugLevel
	}

	file, err := os.OpenFile("out.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error("unable to open log file")
		return f
	}
	f.Context = log.NewEntry(&log.Logger{
		Out:       file,
		Formatter: &log.JSONFormatter{},
		Level:     loglevel,
		Hooks:     make(log.LevelHooks),
	})

	return f
}

func (f File) Debug(format string, args ...any) {
	f.Context.Debugf(format, args...)
}

func (f File) Info(format string, args ...any) {
	f.Context.Infof(format, args...)
}

func (f File) Warn(format string, args ...any) {
	f.Context.Warningf(format, args...)
}

func (f File) Error(format string, args ...any) {
	f.Context.Errorf(format, args...)
}

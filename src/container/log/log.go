package log

import transport "golang-backend-microservice/container/log/transports"

type Transports interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
}

const (
	Console = "console"
	File    = "file"
	Rollbar = "rollbar"
)

var destinations []Transports

func CreateTransports(transports ...string) {
	for _, t := range transports {
		switch t {
		case Console:
			destinations = append(destinations, transport.Console{}.Default())
		case File:
			destinations = append(destinations, transport.File{}.Default())
		case Rollbar:
			destinations = append(destinations, transport.Rollbar{}.Default())
		}
	}
}

func Debug(format string, args ...any) {
	for _, t := range destinations {
		t.Debug(format, args...)
	}
}

func Info(format string, args ...any) {
	for _, t := range destinations {
		t.Info(format, args...)
	}
}

func Warn(format string, args ...any) {
	for _, t := range destinations {
		t.Warn(format, args...)
	}
}

func Error(format string, args ...any) {
	for _, t := range destinations {
		t.Error(format, args...)
	}
}

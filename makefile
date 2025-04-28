NAME=db-backend
VERSION=0.1.0-development

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Welcome to $(shell basename ${PWD})! 🎉"
	@echo
	@echo " Chose a command to run:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /' 
	@echo

## dev: Run package in dev mode
.PHONY: dev
dev:
	go run ./src/cmd/main.go

## start: Test & run package
.PHONY: start
start: test dev

## test: Run package's unit tests
.PHONY: test
test:
	go clean -testcache
	go test -v ./src/container/...

## build: Build package into a Docker image
.PHONY: build
build:
	go build -o ./build/${NAME} ./src/cmd
	docker build -t ${NAME}:${VERSION} ./
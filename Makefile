# load .env file
ifneq (,$(wildcard ./.env))
	include .env
	export $(shell sed 's/=.*//' .env)
endif

.PHONY: default run build test docs clean

default: docs run

run:
	@go run cmd/main.go

build:
	@go build -o $(APP_BINARY_NAME) cmd/main.go

tests:
	@go test ./ ...

docs:
	@swag init -g cmd/main.go

clean:
	@rm -f $(APP_BINARY_NAME)
	@rm -rf docs
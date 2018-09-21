HAVE_GOLINT:=$(shell which golint)
HAVE_MIGRATE:=$(shell which migrate)

## Go
.PHONY: setup lint test build run
setup:
	@echo "Start setup"
	@env GO111MODULE=on go mod vendor

lint: setup golint
	@echo "Check lint"
	@golint $(shell go list ./...|grep -v vendor)
	@go vet ./...

test: setup
	@echo "go test"
	@go test

build: setup
	@echo "build"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/api ./cmd/api

run: setup
	@echo "go run"
	@go run ./cmd/api/main.go -c ./_tools/local/api.toml

## Docker local
CONTAINER_PREFIX:=ddd-go-api-template

.PHONY: dstart dstop dstatus dlogin dclean dlog dmigrate
dstart: setup
	@echo "docker start"
	@docker-compose up -d

dstop:
	@echo "docker stop"
	@docker-compose stop

dstatus:
	@echo "docker status"
	@docker ps --filter name=$(CONTAINER_PREFIX)

dlogin:
	@echo "docker login"
	@docker exec -it $(shell docker ps --all --format "{{.Names}}" | peco) /bin/bash

dclean:
	@echo "docker clean"
	@docker ps --all --filter name=$(CONTAINER_PREFIX) --quiet | xargs docker rm --force

dlog:
	@echo "docker log"
	@docker-compose logs -f $(shell docker ps --all --format "{{.Names}}" | peco | cut -d"_" -f2)

dmigrate: migrate
	@echo "migrate"
	@migrate -path ./_sql -database 'mysql://root:root@tcp(0.0.0.0:3306)/demo' -verbose up

## Install package
.PHONY: golint migrate
golint:
ifndef HAVE_GOLINT
	@echo "Installing linter"
	@go get -u github.com/golang/lint/golint
endif

migrate:
ifndef HAVE_MIGRATE
	@echo "Installing migrate"
	@go get -u -d github.com/mattes/migrate/cli github.com/go-sql-driver/mysql
	@go build -tags 'mysql' -o ${GOPATH}/bin/migrate github.com/mattes/migrate/cli
endif
HAVE_GOLINT:=$(shell which golint)
HAVE_VGO:=$(shell which vgo)
HAVE_MIGRATE:=$(shell which migrate)

## Go
.PHONY: setup lint test build run
setup: vgo
	@echo "Start setup"
	@vgo mod -vendor

lint: setup golint
	@echo "Check lint"
	@golint $(shell go list ./...|grep -v vendor)
	@go vet ./...

test: setup
	@echo "go test"
	@vgo test

build: setup
	@echo "build"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 vgo build -o ./bin/api ./cmd/api

run: setup
	@echo "go run"
	@go run ./cmd/api/main.go -c ./_tools/local/api.toml

## Install package
.PHONY: vgo golint migrate
vgo:
ifndef HAVE_VGO
	@echo "Installing vgo"
	@go get -u golang.org/x/vgo
endif

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
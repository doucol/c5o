REV := $(shell git rev-parse --short HEAD)
SEC := $(shell date +%s)
ifeq ($(shell uname), Darwin)
  DATE := $(shell TZ=UTC date -j -f "%s" ${SEC} +"%Y-%m-%dT%H:%M:%SZ")
else
  DATE := $(shell date -u -d @${SEC} +"%Y-%m-%dT%H:%M:%SZ")
endif

APP := c5o
OUT ?= bin/$(APP)
SRC := github.com/doucol/$(APP)
VER ?= v0.0.1

default: help

test: ## Run tests
	@go clean --testcache && go test ./...

build: ## Build
	@CGO_ENABLED=0 go build -ldflags "-X ${SRC}/cmd.Date=${DATE} -X ${SRC}/cmd.Revision=${REV} -X ${SRC}/cmd.Version=${VER} -w -s" -a -o ${OUT} main.go

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'
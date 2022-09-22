#!/usr/bin/env make

.DEFAULT_GOAL  := help
.DEFAULT_SHELL := /bin/sh

DIST_DIRECTORY := ./dist
BINARY         := $(DIST_DIRECTORY)/gss-api
GIT_COMMIT     := $(shell git rev-parse HEAD)
VERSION        := $(shell git describe --tags --dirty)
LDFLAGS        := "-s -w -X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT)"

DOCKER_REGISTRY            := docker.io
DOCKER_REGISTRY_NAMESPACE  := jjuarez
DOCKER_REGISTRY_IMAGE_NAME := gss-api
DOCKER_IMAGE_NAME          := $(DOCKER_REGISTRY)/$(DOCKER_REGISTRY_NAMESPACE)/$(DOCKER_REGISTRY_IMAGE_NAME)

.PHONY: help
help: ## Shows this pretty help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z//_-]+:.*?##/ { printf " %-20s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

$(BINARY):
	@go build -o $(BINARY) -v main.go

.PHONY: build
build: $(BINARY) ## Build the service binary

.PHONY: clean
clean: ## Clean all the go temporary stuff
	@rm -f $(BINARY)

.PHONY: docker/build
docker/build: ## Creates a docker image for the service
	@docker image build --tag $(DOCKER_IMAGE_NAME):latest --file Dockerfile .

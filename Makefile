#!/usr/bin/env make

.DEFAULT_GOAL  := help
.DEFAULT_SHELL := /bin/bash

DIST_DIRECTORY := ./bin
GIT_COMMIT     := $(shell git rev-parse --verify HEAD 2>/dev/null)
VERSION        := $(shell git describe --tags 2>/dev/null)

DOCKER_REGISTRY            := docker.io
DOCKER_REGISTRY_NAMESPACE  := jjuarez
DOCKER_REGISTRY_IMAGE_NAME := gss-api
DOCKER_IMAGE_NAME          := $(DOCKER_REGISTRY)/$(DOCKER_REGISTRY_NAMESPACE)/$(DOCKER_REGISTRY_IMAGE_NAME)

.PHONY: help
help: ## Shows this pretty help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z//_-]+:.*?##/ { printf " %-20s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## Clean all the go temporary stuff
	@go clean
	@rm -f $(DIST_DIRECTORY)/*

.PHONY: build
build: ## Build the service binary
ifdef VERSION
	@go build -ldflags "-X 'main.GitCommit=$(GIT_COMMIT)' -X 'main.Version=$(VERSION)'" -o $(DIST_DIRECTORY) -v ./...
else
	@go build -ldflags "-X 'main.GitCommit=$(GIT_COMMIT)' -X 'main.Version=v0.0.0'" -o $(DIST_DIRECTORY) -v ./...
endif

.PHONY: run
run: ## Run the service
	@$(DIST_DIRECTORY)/api

.PHONY: docker/login
docker/login:
ifdef DOCKER_PASSWORD
	@echo $(DOCKER_PASSWORD) | docker login --username $(DOCKER_USERNAME) --password-stdin $(DOCKER_REGISTRY)
else
	$(warning "The Docker login needs authentication data, DOCKER_USERNAME, and DOCKER_PASSWORD")
endif

.PHONY: docker/build
docker/build: ## Creates a docker image for the service
ifdef VERSION
	@docker image build \
		--cache-from $(DOCKER_IMAGE_NAME):latest \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		--build-arg VERSION=$(VERSION) \
		--tag $(DOCKER_IMAGE_NAME):$(GIT_COMMIT) \
		--tag $(DOCKER_IMAGE_NAME):$(VERSION) \
		--tag $(DOCKER_IMAGE_NAME):latest \
		--target runtime \
		--file Dockerfile \
		.
else
	@docker image build \
		--cache-from $(DOCKER_IMAGE_NAME):latest \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		--build-arg VERSION=v0.0.0 \
		--tag $(DOCKER_IMAGE_NAME):$(GIT_COMMIT) \
		--tag $(DOCKER_IMAGE_NAME):latest \
		--target runtime \
		--file Dockerfile \
		.
endif

.PHONY: docker/release
docker/release: docker/build docker/login ## Pubishes a docker release
ifdef VERSION
		docker image push $(DOCKER_IMAGE_NAME):$(VERSION)
else
	$(warning "To generate a valid release you a valid tag")
endif

.DEFAULT_GOAL := help
PROJECT = healthChecker
BUILD 	= $(PWD)/build

# `make help` generates a help message for each target that
# has a comment starting with ##
help:
	@echo "Please use 'make <target>' where <target> is one of the following:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	GOOS=$(goos) GO111MODULE=on go build -o $(BUILD)/$(PROJECT) ./cmd/$(PROJECT)/main.go

.PHONY: help build
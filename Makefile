.DEFAULT_GOAL := help
.PHONY: help clean build run docker-build docker-run docker-stop

# Define variables
BINARY_NAME := go-app
DOCKER_IMAGE := make-go-app
DOCKER_CONTAINER := make-go-app
MAIN_FILE := cmd/app/main.go
LDFLAGS := -ldflags="-s -w"

help: ## Display this help message
	@echo "Usage: make <command>"
	@echo ""
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

clean: ## Remove binary file and container resources
	rm -f $(BINARY_NAME)
	docker stop $(DOCKER_CONTAINER) || true
	docker rm $(DOCKER_CONTAINER) || true
	docker rmi $(DOCKER_IMAGE) || true

build: ## Build the binary file
	go build $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_FILE)

run: ## Build and run the binary file
	go run $(LDFLAGS) $(MAIN_FILE)

docker-build: ## Build Docker image
	docker build -t $(DOCKER_IMAGE) . -f ./Containerfile

docker-run: ## Run Docker container
	@if ! docker image inspect $(DOCKER_IMAGE) >/dev/null 2>&1; then \
		make docker-build; \
	fi
	docker run -d -p 8080:8080 --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)


docker-stop: ## Stop Docker container
	docker stop $(DOCKER_CONTAINER)
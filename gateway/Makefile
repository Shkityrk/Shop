.PHONY: build run docker-build docker-run swagger clean help

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	go build -o gateway .

run: ## Run the application
	go run main.go

swagger: ## Generate swagger documentation
	swag init

docker-build: ## Build docker image
	docker build -t gateway:latest .

docker-run: ## Run docker container
	docker run -p 8000:8000 --name gateway-container gateway:latest

clean: ## Clean build artifacts
	rm -f gateway
	go clean

test: ## Run tests
	go test -v ./...

deps: ## Download dependencies
	go mod download
	go mod tidy

lint: ## Run linter
	golangci-lint run


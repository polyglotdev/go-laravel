BINARY_NAME=celeritasApp

build: ## Build the application binary
	@go mod vendor
	@echo "Building Celeritas..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Celeritas built!"

run: build ## Build and run the application
	@echo "Starting Celeritas..."
	@./tmp/${BINARY_NAME} &
	@echo "Celeritas started!"

clean: ## Clean up the binary
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test: ## Run all tests in any directory
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run ## Start the application

stop: ## Stop the application
	@echo "Stopping Celeritas..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Celeritas!"

restart: stop start ## Restart the application

.PHONY: build run clean test start stop restart

help: ## Display details on all commands
	@awk 'BEGIN {FS = ":.*?##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n%s\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

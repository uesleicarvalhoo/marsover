COVERAGE_OUTPUT=coverage.output
COVERAGE_HTML=coverage.html

GO ?= go
ENV_FILE ?= .env
-include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

## @ Help
.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_/-]+:.*?##/ { printf "\033[36m%-18s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

## @ Tools
.PHONY: install-tools
install-tools:  ## Install gofumpt, gocritic and swaggo
	@$(GO) install mvdan.cc/gofumpt@latest
	@$(GO) install github.com/go-critic/go-critic/cmd/gocritic@latest
	@$(GO) install github.com/swaggo/swag/cmd/swag@latest

## @ Linter
.PHONY: lint format
lint:  ## Run golangci-lint and gocritic
	@golangci-lint run ./...
	@gocritic check ./...
	@golangci-lint run ./... --exclude-dirs=docs

format:  ## Format code
	@gofumpt -e -l -w .

## @ Tests
.PHONY: test test/unit test/integration coverage
test:  ## Run all tests
	@$(GO) test -covermode=atomic -count=1 -v ./... -p 1 -race -coverprofile $(COVERAGE_OUTPUT) | gotestfmt

test/unit:  ## Run unit tests only
	@$(GO) test -covermode=atomic -short -v ./... -p 1 -race

test/coverage: ## Run tests and generate coverage report
	@$(GO) test ./... -covermode=atomic -count=1 -race -coverprofile $(COVERAGE_OUTPUT)
	@$(GO) tool cover -html=$(COVERAGE_OUTPUT) -o $(COVERAGE_HTML)

test/coverage-browser: test/coverage ## Open coverage report in browser
	@xdg-open $(COVERAGE_HTML) || open $(COVERAGE_HTML) || wslview $(COVERAGE_HTML)

## @ Application
.PHONY: swagger run
swagger: ## Generate Swagger docs
	@echo "📝 Generating Swagger docs..."
	@swag init --generalInfo ./internal/http/server.go --output ./docs

run: ## Run HTTP server
	@$(GO) run main.go

## @ Docker
.PHONY: docker compose-up compose-down
docker: ## Build Docker image locally
	@docker build -t marsrover:latest .

compose/up: ## Start services with docker-compose
	@docker compose --env-file $(ENV_FILE) up -d --build

compose/down: ## Stop services
	@docker-compose --env-file $(ENV_FILE) down

## @ Clean
.PHONY: clean clean_coverage_cache
clean_coverage_cache:
	@rm -rf $(COVERAGE_OUTPUT) $(COVERAGE_HTML)

clean: clean_coverage_cache ## Remove cache files

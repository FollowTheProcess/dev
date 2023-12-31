.PHONY: help tidy fmt test bench lint cover clean check sloc install uninstall
.DEFAULT_GOAL := help

COVERAGE_DATA := coverage.out
COVERAGE_HTML := coverage.html

export GOEXPERIMENT := loopvar

help: ## Show the list of available tasks
	@echo "Available Tasks:\n"
	@grep -E '^[a-zA-Z_0-9%-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-10s %s\n", $$1, $$2}'

tidy: ## Tidy dependencies in go.mod
	go mod tidy

fmt: ## Run go fmt on all source files
	go fmt ./...

test: ## Run the test suite
	go test -race ./...

bench: ## Run benchmarks
	go test ./... -run=None -benchmem -bench .

lint: ## Run the linters and auto-fix if possible
	golangci-lint run --fix

cover: ## Calculate test coverage and render the html
	go test -race -cover -covermode atomic -coverprofile $(COVERAGE_DATA) ./...
	go tool cover -html $(COVERAGE_DATA) -o $(COVERAGE_HTML)
	open $(COVERAGE_HTML)

clean: ## Remove build artifacts and other clutter
	go clean ./...
	rm -rf $(COVERAGE_DATA) $(COVERAGE_HTML) ./bin ./dist

check: test lint ## Run tests and linting in one go

sloc: ## Print lines of code (for fun)
	find . -name "*.go" | xargs wc -l | sort -nr | head

build: ## Compile the project binary
	mkdir -p ./bin
	goreleaser build --single-target --skip before --snapshot --clean --output ./bin/dev

install: uninstall build ## Install the project on your machine
	cp ./bin/dev ${GOBIN}

uninstall: ## Uninstall the project from your machine
	rm -rf ${GOBIN}/dev

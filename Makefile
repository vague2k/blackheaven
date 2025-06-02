all: lint format test

.PHONY: lint
lint:
	@echo "Running linter...";
	@golangci-lint run --color=always;

.PHONY: format
format:
	@echo "Running formatter...";
	@gofumpt -l -w -d .;

.PHONY: test
test:
	@echo "Running tests...";
	@set -euo pipefail
	@go test -json $(go list ./... | grep -v '/ui') | tee /tmp/gotest.log | gotestfmt -hide=empty-packages 
templ:
	@templ generate --watch --proxy="localhost:3000" --open-browser=false

server:
	@air

tailwind:
	@tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

dev:
	@make -j3 tailwind templ server

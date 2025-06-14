lint:
	@echo "Running linter...";
	@golangci-lint run --color=always;

# uncomment this when gofumpt implements ignore mechanism for go tooling, 
# see https://github.com/golang/go/issues/42965 for more details
format:
# @echo "Running formatter...";
# @gofumpt -l -w -d .;

test:
	@echo "Running tests...";
	@set -euo pipefail
	@go test -json $(go list ./... | grep -v '/internal') | tee /tmp/gotest.log | gotestfmt -hide=empty-packages 

templ:
	@templ generate --watch --proxy="localhost:3000" --open-browser=false

server:
	@air

tailwind:
	@tailwindcss -i ./internal/assets/css/input.css -o ./internal/assets/css/output.css --watch

dev:
	@make -j3 tailwind templ server

components:
	@templui add selectbox form input label textarea toast

vet: lint format test 

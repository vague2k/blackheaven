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

air:
	@air \
    --build.cmd "go build -o tmp/bin/main ./main.go" \
    --build.bin "tmp/bin/main" \
    --build.delay "100" \
    --build.exclude_dir "node_modules" \
    --build.include_ext "go" \
    --build.stop_on_error "false" \
    --misc.clean_on_exit true

tailwind:
	@tailwindcss -i ./views/assets/css/input.css -o ./views/assets/css/output.css --watch

dev:
	@make -j3 tailwind templ air

components:
	@templui add selectbox form input label textarea toast

vet: lint format test

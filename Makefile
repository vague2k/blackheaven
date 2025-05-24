# dont mind this makefile, I like pretty outputs :)
GREEN = \033[32m
BLUE = \033[34m
RESET = \033[0m

all: lint format test

.PHONY: lint
lint:
	@echo "$(BLUE)::$(RESET) Running golangci-lint";
	@golangci-lint run --color=always;

.PHONY: format
format:
	@echo "$(BLUE)::$(RESET) Running gofumpt formatter";
	@gofumpt -l -w -d .;

.PHONY: test
test:
	@echo "$(BLUE)::$(RESET) Running tests";
# See https://github.com/gotestyourself/gotestsum
	@gotestsum --format-hide-empty-pkg; 

.PHONY: hot-tests
hot-tests:
	@echo "$(BLUE)::$(RESET) Running tests";
# See https://github.com/gotestyourself/gotestsum
	@gotestsum --watch --format-hide-empty-pkg ./...; 

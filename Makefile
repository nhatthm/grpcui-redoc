MODULE_NAME=grpcui-redoc

VENDOR_DIR = vendor
NODE_MODULES_DIR = resources/web/node_modules
CSS_FILE ?= css.go

GOLANGCI_LINT_VERSION ?= v1.58.0

GO ?= go
GOLANGCI_LINT ?= $(shell go env GOPATH)/bin/golangci-lint-$(GOLANGCI_LINT_VERSION)

GITHUB_OUTPUT ?= /dev/null

.PHONY: $(VENDOR_DIR)
$(VENDOR_DIR):
	@mkdir -p $(VENDOR_DIR)
	@$(GO) mod vendor
	@$(GO) mod tidy

.PHONY: $(NODE_MODULES_DIR)
$(NODE_MODULES_DIR):
	@cd resources/web; npm install --save-dev stylelint stylelint-config-standard postcss postcss-cli autoprefixer

.PHONY: lint
lint: lint-go lint-css

.PHONY: lint-go
lint-go: $(GOLANGCI_LINT)
	@echo ">> golang lint"
	@$(GOLANGCI_LINT) run -c .golangci.yaml

.PHONY: lint-css
lint-css: $(NODE_MODULES_DIR)
	@echo ">> css lint"
	@cd resources/web; npx stylelint "**/*.css"

.PHONY: generate
generate: $(NODE_MODULES_DIR)
	@resources/scripts/generate.sh "$(CSS_FILE)"

.PHONY: cleanup
cleanup: cleanup-css
	@rm -rf $(VENDOR_DIR)

.PHONY: cleanup-css
cleanup-css:
	@rm -rf $(NODE_MODULES_DIR) resources/web/package.json resources/web/package-lock.json

.PHONY: test
test: test-unit test-css

## Run unit tests
.PHONY: test-unit
test-unit:
	@echo ">> unit test"
	@$(GO) test -gcflags=-l -coverprofile=unit.coverprofile -covermode=atomic -race ./... -tags testcoverage

.PHONY: test-css
test-css: $(NODE_MODULES_DIR)
	@echo ">> css test"
	@resources/scripts/test.sh

.PHONY: $(GITHUB_OUTPUT)
$(GITHUB_OUTPUT):
	@echo "MODULE_NAME=$(MODULE_NAME)" >> "$@"
	@echo "GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION)" >> "$@"

$(GOLANGCI_LINT):
	@echo "$(OK_COLOR)==> Installing golangci-lint $(GOLANGCI_LINT_VERSION)$(NO_COLOR)"; \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin "$(GOLANGCI_LINT_VERSION)"
	@mv ./bin/golangci-lint $(GOLANGCI_LINT)

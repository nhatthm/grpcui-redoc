VENDOR_DIR = vendor
NODE_MODULES_DIR = resources/web/node_modules

CSS_FILE ?= css.go

GO ?= go
GOLANGCI_LINT ?= golangci-lint

.PHONY: $(VENDOR_DIR) lint lint-go lint-css test test-unit test-css

$(VENDOR_DIR):
	@mkdir -p $(VENDOR_DIR)
	@$(GO) mod vendor
	@$(GO) mod tidy

$(NODE_MODULES_DIR):
	@cd resources/web; npm install --save-dev stylelint stylelint-config-standard postcss postcss-cli autoprefixer

lint: lint-go lint-css

lint-go:
	@echo ">> golang lint"
	@$(GOLANGCI_LINT) run

lint-css: $(NODE_MODULES_DIR)
	@echo ">> css lint"
	@cd resources/web; npx stylelint "**/*.css"

generate: $(NODE_MODULES_DIR)
	@resources/scripts/generate.sh "$(CSS_FILE)"

cleanup: cleanup-css
	@rm -rf $(VENDOR_DIR)

cleanup-css:
	@rm -rf $(NODE_MODULES_DIR) resources/web/package.json resources/web/package-lock.json

test: test-unit test-css

test-unit:
	@echo ">> unit test"
	@$(GO) test -gcflags=-l -coverprofile=unit.coverprofile -covermode=atomic -race ./... -tags testcoverage

test-css: $(NODE_MODULES_DIR)
	@echo ">> css test"
	@resources/scripts/test.sh

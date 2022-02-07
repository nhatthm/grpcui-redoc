VENDOR_DIR = vendor

GO ?= go
GOLANGCI_LINT ?= golangci-lint

.PHONY: $(VENDOR_DIR) lint test test-unit

$(VENDOR_DIR):
	@mkdir -p $(VENDOR_DIR)
	@$(GO) mod vendor
	@$(GO) mod tidy

lint:
	@$(GOLANGCI_LINT) run

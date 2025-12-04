BIN_DIR := ./bin
BIN_PATH := $(BIN_DIR)/time_tracker_backend
CMD_DIR := ./cmd/backend
OPENAPI_SPEC_PATH := ../api/openapi.yaml
OAPI_CODEGEN_CONFIG_PATH := ./oapi_codegen_config.yaml

.PHONY: run
run:
	go run $(CMD_DIR)

.PHONY: build
build:
	go build -o $(BIN_PATH) $(CMD_DIR)

.PHONY: dependencies
dependencies:
	go mod tidy
	go mod vendor

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: generate-api
generate-api:
	go tool oapi-codegen -config $(OAPI_CODEGEN_CONFIG_PATH) $(OPENAPI_SPEC_PATH)

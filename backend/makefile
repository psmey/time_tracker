OPENAPI_SPEC_PATH := ../openapi.yaml
OAPI_CODEGEN_CONFIG_PATH := ./oapi_codegen_config.yaml

.PHONY: generate-api
generate-api:
	go tool oapi-codegen -config $(OAPI_CODEGEN_CONFIG_PATH) $(OPENAPI_SPEC_PATH)

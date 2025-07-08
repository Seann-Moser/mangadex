SWAGGER_URL=https://api.mangadex.org/docs/static/api.yaml
SWAGGER_FILE=api.yaml
OUTPUT_DIR=./

.PHONY: all clean generate

all: generate

download:
	curl -L $(SWAGGER_URL) -o $(SWAGGER_FILE)

generate: download
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	oapi-codegen --config=model-cfg.yaml api.yaml
	oapi-codegen --config=cfg.yaml api.yaml
	go run cache_create/main.go -input client.gen.go -interface ClientWithResponsesInterface -output client_cache.go   -package mangadex
clean:
	rm -rf $(OUTPUT_DIR) $(SWAGGER_FILE)

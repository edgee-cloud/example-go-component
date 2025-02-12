.PHONY: all
MAKEFLAGS += --silent

all: help

help:
	@grep -E '^[a-zA-Z1-9\._-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| sed -e "s/^Makefile://" -e "s///" \
		| awk 'BEGIN { FS = ":.*?## " }; { printf "\033[36m%-30s\033[0m %s\n", $$1, $$2 }'
internal:
	go run go.bytecodealliance.org/cmd/wit-bindgen-go generate -o internal/ ./wit
setup: internal ## setup development environment

build:
	edgee components build

build-no-edgee: setup
	tinygo build -target=wasip2 -o dc_component.wasm --wit-package wit/ --wit-world data-collection ./

clean: ## clean build artifacts
	rm -rf dc_component.wasm
	rm -rf internal/

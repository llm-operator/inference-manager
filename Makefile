.PHONY: default
default: test

include common.mk

.PHONY: test
test: go-test-all

.PHONY: lint
lint: go-lint-all git-clean-check

.PHONY: generate
generate: buf-generate-all

.PHONY: build-server
build-server:
	go build -o ./bin/server ./server/cmd/

.PHONY: build-engine
build-engine:
	go build -o ./bin/engine ./engine/cmd/

.PHONY: build-docker-server
build-docker-server:
	docker build --build-arg TARGETARCH=amd64 -t llm-operator/inference-manager-server:latest -f build/server/Dockerfile .

.PHONY: build-docker-engine
build-docker-engine:
	docker build --build-arg TARGETARCH=amd64 --target ollama -t llm-operator/inference-manager-engine:latest -f build/engine/Dockerfile .

.PHONY: build-docker-engine-vllm
build-docker-engine-vllm:
	docker build --build-arg TARGETARCH=amd64 --target vllm -t llm-operator/inference-manager-engine:vllm_latest -f build/engine/Dockerfile .

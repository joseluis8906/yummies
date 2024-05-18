BAZEL_BIN = $$(bazel info bazel-bin)
PROJECT_NAME = yummies

.PHONY: build
build: 
	@go build -ldflags "-s -w" -o bin/release/$(PROJECT_NAME) cmd/$(PROJECT_NAME)/main.go


.PHONY: test
test:
	@go test -cov ./...

.PHONY: run
run:
	@CONFIG_URL="http://yummies.local:2379" OTEL_SERVICE_NAME="delivery" ./bin/release/$(PROJECT_NAME)

.PHONY: debug
debug:
	@go build -gcflags=all="-N -l" -o ./bin/debug/$(PROJECT_NAME) cmd/$(PROJECT_NAME)/main.go
	@dlv debug --headless --listen=:2345 --api-version=2 -- ./bin/debug/$(PROJECT_NAME)

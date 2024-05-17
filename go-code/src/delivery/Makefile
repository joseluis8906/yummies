BAZEL_BIN = $$(bazel info bazel-bin)
PROJECT_NAME = delivery

.PHONY: build
build: pb
	@go build -ldflags "-s -w" -o bin/release/$(PROJECT_NAME) cmd/delivery/main.go

.PHONY: pb
pb:
	@protoc --proto_path=../../protobuf/yummies\
		--go_out=./pkg/pb --go_opt=paths=source_relative\
		--go-grpc_out=./pkg/pb --go-grpc_opt=paths=source_relative\
		../../protobuf/yummies/home.proto

.PHONY: test
test:
	@go test -cov ./...

.PHONY: run
run:
	@CONFIG_URL="http://delivery.local:2379" OTEL_SERVICE_NAME="delivery" ./bin/release/$(PROJECT_NAME)

.PHONY: debug
debug:
	@go build -gcflags=all="-N -l" -o ./bin/debug/$(PROJECT_NAME) cmd/delivery/main.go
	@dlv debug --headless --listen=:2345 --api-version=2 -- ./bin/debug/$(PROJECT_NAME)


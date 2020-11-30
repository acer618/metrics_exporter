PACKAGE_NAME=lightstep-exporter
VERSION=0.1.0

SRC_DIR=src
GITHUB_DIR=src/github.com
OTEL_DIR:=$(GITHUB_DIR)/open-telemetry
OTEL_PROTOS:=$(OTEL_DIR)/opentelemetry/proto/metrics/v1/metrics.proto $(OTEL_DIR)/opentelemetry/proto/collector/metrics/v1/metrics_service.proto $(OTEL_DIR)/opentelemetry/proto/common/v1/common.proto $(OTEL_DIR)/opentelemetry/proto/resource/v1/resource.proto

export GOPATH := $(shell pwd)

deps:
	@echo Getting dependencies...
	@go get google.golang.org/grpc
	@go get github.com/golang/protobuf/protoc-gen-go

otel: 
	@if [ ! -d "$(GITHUB_DIR)/open-telemetry" ] ;\
	then \
		git clone https://github.com/open-telemetry/opentelemetry-proto.git $(OTEL_DIR); \
	fi

.PHONY: protobuf
protobuf: deps otel $(OTEL_PROTOS)
	@sudo protoc --proto_path=$(OTEL_DIR) --plugin=bin/protoc-gen-go --go_out=plugins=grpc:$(SRC_DIR) $(OTEL_PROTOS)

.PHONY: protobuf build
build:
	go build -ldflags '-X main.version=$(VERSION)' -o bin/metrics_client src/main.go
	go build -ldflags '-X main.version=$(VERSION)' -o bin/metrics_server src/metrics_service.go src/metrics_server.go

.PHONY: test
test: build

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: clean
clean:
	sudo rm -rf src/github.com

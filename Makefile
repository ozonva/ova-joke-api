LOCAL_BIN:=$(CURDIR)/bin

run:
	go run cmd/ova-joke-api/main.go

lint:
	golangci-lint run ./...

test:
	go test -tags=test_unit -v -count=1 -race -timeout=1m ./...

.PHONY: build
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ova-joke-api
		PATH=$(LOCAL_BIN):$(PATH) protoc -I vendor.protogen \
				--go_out=pkg/ova-joke-api --go_opt=paths=import \
				--go-grpc_out=pkg/ova-joke-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ova-joke-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ova-joke-api/ova-joke-api.proto
		mv pkg/ova-joke-api/github.com/ozonva/ova-joke-api/pkg/ova-joke-api/* pkg/ova-joke-api/
		rm -rf pkg/ova-joke-api/gitlab.com
		mkdir -p cmd/ova-joke-api
		cd pkg/ova-joke-api && ls go.mod || go mod init github.com/ozonva/ova-joke-api/pkg/ova-joke-api && go mod tidy

.PHONY: generate
generate: .vendor-proto .generate

.PHONY: build
build:
		go build -o $(LOCAL_BIN)/ova-joke-api cmd/ova-joke-api/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ova-joke-api
		cp api/ova-joke-api/ova-joke-api.proto vendor.protogen/api/ova-joke-api/ova-joke-api.proto
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init github.com/ozonva/ova-joke-api
		GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
		GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
		GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
		GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger


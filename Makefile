check: test vet fmt lint

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
		test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
	golangci-lint run

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/proto/hash.go.proto

build:
	go build -o bin/refresh_hash_worker ./cmd/refresh_hash_worker
	go build -o bin/http_api ./cmd/http_api
	go build -o bin/grpc_api ./cmd/grpc_api
	go build -o bin/grpc_client ./cmd/grpc_client

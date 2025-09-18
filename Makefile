.PHONY: run build test clean

run:
	go run cmd/gateway/main.go

build:
	go build -o bin/gateway cmd/gateway/main.go

test:
	go test -v ./...

clean:
	rm -f bin/gateway

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

dev:
	./scripts/run-dev.sh
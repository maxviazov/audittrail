VERSION := $(shell cat VERSION)

.PHONY: all version proto up run lint test clean

all: build

version:
	@echo "📦 Version: $(VERSION)"

proto:
	@echo "🔧 Generating gRPC files..."
	buf generate

up:
	docker-compose up --build

run:
	go run ./cmd/server

lint:
	golangci-lint run ./...

test:
	go test ./... -v

clean:
	rm -rf bin/ dist/
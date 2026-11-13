.PHONY: all build test vet

all: build

build:
	mkdir -p dist
	go build -trimpath -o dist/xynaut ./cmd/xynaut

test:
	go test ./...

vet:
	go vet ./...

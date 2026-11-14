.PHONY: all build test vet

all: build

build:
	mkdir -p dist
	go build -trimpath -o dist/xynaut ./cmd/xynaut

test:
	@test -f protocol/specs/src/control/methods.json || { \
	  echo 'Protocol specs are missing. Clone with --recurse-submodules or run: git submodule update --init --recursive' >&2; \
	  exit 1; \
	}
	go test ./...

vet:
	go vet ./...

BIN ?= xpo

hello:
	@echo "Hello, world!"

build:
	@echo "Building bin/$(BIN)..."
	@go build -ldflags="-s -w" -o bin/$(BIN) ./cmd/xpo.go


test:
	@go test `go list ./... | grep -v cmd`

clean:
	@echo "Removing bin/$(BIN)..."
	@rm -f bin/$(BIN)

all: clean build

.DEFAULT_GOAL := all
.PHONY: all format install clean upgrade test

BIN = bin
CFLAGS = -ldflags "-w -s" -x
PKG = ./cmd/uast

all:
	@echo "=====Building locally====="
	go build -v -o $(BIN)/ $(CFLAGS) $(PKG)

format:
	gofmt -s -w $(shell find . -name "*.go")

install: upgrade all
	@echo "=====Installing binary to PATH====="
	go install $(CFLAGS) -v $(PKG)

clean:
	go clean -i -x $(PKG)
	rm -rfv $(BIN)

upgrade:
	@echo "=====Upgrading dependencies====="
	go get -u -v ./...
	go mod tidy

test:
	go test -v -cover ./...

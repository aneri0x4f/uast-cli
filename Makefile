.DEFAULT_GOAL := all
.PHONY: all format install clean upgrade test

BIN = bin/
CFLAGS = -ldflags "-s" -x -v -trimpath -gcflags "-m"
PKG = ./cmd/uast

all:
	@echo "=====Building locally====="
	GOEXPERIMENT=greenteagc go build -o $(BIN) $(CFLAGS) $(PKG)

format:
	gofmt -e -s -w $(shell find . -name "*.go")

install: upgrade all
	@echo "=====Installing binary to PATH====="
	GOEXPERIMENT=greenteagc go install $(CFLAGS) $(PKG)

clean:
	go clean -i -x $(PKG)
	rm -rfv $(BIN)

upgrade:
	@echo "=====Upgrading dependencies====="
	go get -u -v ./...
	go mod tidy

test:
	go test -v -cover ./...

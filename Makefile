.DEFAULT_GOAL := all
.PHONY: all format install clean upgrade test

BIN = bin
CFLAGS = -ldflags "-w -s" -x
PKG = ./cmd/uast

all:
	go build -v -o $(BIN)/ $(CFLAGS) $(PKG)

format:
	gofmt -s -w $(shell find . -name "*.go")

install: upgrade
	go install $(CFLAGS) -v $(PKG)

clean:
	go clean -i -x $(PKG)
	rm -rfv $(BIN)

upgrade:
	go get -u -v ./...
	go mod tidy

test:
	go test -v -cover ./...

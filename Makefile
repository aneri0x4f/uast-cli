.DEFAULT_GOAL := all
.PHONY: all docker podman format install clean upgrade test

BIN = bin/uast
CFLAGS = -ldflags "-w -s" -x
GOPATH = $(shell go env GOPATH)
PKG = ./cmd/uast

all:
	go build -v -o $(BIN) $(CFLAGS) $(PKG)

format:
	gofmt -s -w **/*.go

install: upgrade
	go install $(PKG)

clean:
	go clean -i -n
	rm -rfv $(BIN) $(GOPATH)/bin/uast

upgrade:
	go get -u -v ./...
	go mod tidy

test:
	go test -v -cover ./...

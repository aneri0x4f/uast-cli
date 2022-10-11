.DEFAULT_GOAL := all
.PHONY: all docker format install clean upgrade test

BIN = bin/uast
CFLAGS = -ldflags "-w -s" -x
GOPATH = $(shell go env GOPATH)

all:
	go build -v -o $(BIN) $(CFLAGS) .

podman:
	podman build --pull --compress -t uast .
	podman system prune -f

docker:
	docker buildx build --pull --compress -t uast .
	docker system prune -f

format:
	gofmt -s -w **/*.go

install: upgrade all
	cp $(BIN) $(GOPATH)/bin

clean:
	go clean -i -n
	rm -rfv $(BIN) $(GOPATH)/bin/uast

upgrade:
	go get -u -v
	go mod tidy

test:
	go test -v -cover ./...

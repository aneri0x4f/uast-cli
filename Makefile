.DEFAULT_GOAL := all
.PHONY: all docker format install clean

CFLAGS = -ldflags "-w -s"
BIN = bin/uast

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

install: all
	cp bin/uast $(shell go env GOPATH)/bin

clean:
	go clean -i -n
	rm -rfv $(BIN) $(shell go env GOPATH)/bin/uast

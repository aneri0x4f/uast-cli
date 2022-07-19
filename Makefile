.DEFAULT_GOAL := all
.PHONY: all docker format install clean

CFLAGS = -ldflags "-w -s"
BIN = bin/uast

all:
	go build -v -o $(BIN) $(CFLAGS) .

docker:
	docker buildx build --pull --compress -t uast .
	docker system prune -f

format:
	gofmt -s -w **/*.go

install:
	go install -v $(CFLAGS)

clean:
	go clean -i -n
	rm -rfv $(BIN)

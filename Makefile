.DEFAULT_GOAL := all

CFLAGS = -ldflags "-w -s"
BIN = bin/

all:
	go build -v -o $(BIN) $(CFLAGS) .

docker:
	docker build --pull --compress -t uast .
	docker system prune -f

install:
	go install -v $(CFLAGS)

clean:
	go clean -i -n
	rm -rfv $(BIN)

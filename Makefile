.DEFAULT_GOAL := all

CFLAGS = -ldflags "-w -s"
BIN = bin/

all:
	go build -o $(BIN) $(CFLAGS) .

install:
	go install $(CFLAGS)

clean:
	go clean -i -n
	rm -rf $(BIN)

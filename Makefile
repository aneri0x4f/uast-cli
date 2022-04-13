.DEFAULT_GOAL := all

CFLAGS = -ldflags "-w -s"
BIN = bin/

all:
	go build -v -o $(BIN) $(CFLAGS) .

install:
	go install -v $(CFLAGS)

clean:
	go clean -i -n
	rm -rfv $(BIN)

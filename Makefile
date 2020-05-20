BIN := $(shell basename $(CURDIR))
VERSION := $(shell git describe --abbrev=0 --tags || echo 'v0.0.0')
REVISION := $(shell git rev-list -1 HEAD || echo 'NONE')

.PHONY: all clean

all: clean test $(BIN)

clean:
	rm -f $(BIN)
	go clean .

test:
	go test -v .

$(BIN):
	go build -o $@ -ldflags '-X github.com/musl/chaat/cmd.Version=$(VERSION) -X github.com/musl/chaat/cmd.Revision=$(REVISION)' .

.PHONY: all clean tools


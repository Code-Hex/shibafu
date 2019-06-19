VERSION := $(shell git rev-parse HEAD | cut -c 1-7)

LDFLAGS = -ldflags="-X main.version=${VERSION}"

.PHONY: build
build:
	go build $(LDFLAGS) -o bin/shibafu github.com/Code-Hex/shibafu/cmd/shibafu
GO_FILES=$(shell find . -name '*.go')

build:
	go build -o bin/intcast ./cmd/intcast

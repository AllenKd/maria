.PHONY: build all

build:
	go build -o maria ./cmd/main

build-linux:
	GOOS=linux GOARCH=amd64 go build -o maria ./cmd/main

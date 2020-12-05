.PHONY: build
build:
	go build -o server -v ./cmd 
test:
	go test -v ./...
integration-test:
	go test -v ./...
.DEFAULT_GOAL := build
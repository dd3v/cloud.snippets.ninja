.PHONY: build
build:
	go build -o server -v ./cmd 
.DEFAULT_GOAL := build
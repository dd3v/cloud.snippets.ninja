.PHONY: build
build:
	go build -o server -v ./cmd 
test:
	go test -v ./...
integration-test:
	go test -v ./...
migrate:
	migrate  -path migrations -database "mysql://root:root@tcp(localhost:3306)/snippets" -verbose up	
.DEFAULT_GOAL := build
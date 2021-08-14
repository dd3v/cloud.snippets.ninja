BINARY_NAME=server
MODULE = $(shell go list -m)
CONFIG_FILE ?= ./config/local.yml
DATABASE_DNS ?= $(shell sed -n 's/^migration_db_dns:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
.PHONY: build
build: ## build the app bin
	CGO_ENABLED=0 go build -o ./${BINARY_NAME} $(MODULE)/cmd/server
.PHONY: test
test:
	go test -v ./...
.PHONY: integration-test
integration-test:
	go test -v ./...
.PHONY: migrate
migrate:
	migrate -path migrations -database "${DATABASE_DNS}" -verbose up
.PHONY: migrate-create
migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -seq -dir ./migrations/ $${name// /_}
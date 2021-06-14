.PHONY: build
build:
	go build -o server -v ./cmd
.PHONY: test
test:
	go test -v ./...
.PHONY: integration-test
integration-test:
	go test -v ./...
.PHONY: migrate
migrate:
	migrate -path migrations -database "mysql://root:root@tcp(localhost:3306)/snippets" -verbose up
.PHONY: migrate-create
migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -seq -dir ./migrations/ $${name// /_}
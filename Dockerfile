FROM golang:alpine AS build
RUN go version
RUN apk update
RUN apk --no-cache add alpine-sdk curl git bash make ca-certificates
RUN rm -rf /var/cache/apk/*
ARG MIGRATE_VERSION=4.7.1
ADD https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz /tmp
RUN tar -xzf /tmp/migrate.linux-amd64.tar.gz -C /usr/local/bin && mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate
WORKDIR /app
RUN pwd
COPY go.* ./
RUN go mod download
RUN go mod verify
COPY .. .
RUN ls -all
RUN go list -m
RUN make test
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
RUN mkdir -p /var/log/app
WORKDIR /app
COPY --from=build /usr/local/bin/migrate /usr/local/bin
COPY --from=build /app/migrations ./migrations/
COPY --from=build /app/config/*.yml ./config/
COPY --from=build /app/server .
COPY --from=build /app/entrypoint.sh .
RUN ls -la
ENTRYPOINT ["bash", "./entrypoint.sh"]
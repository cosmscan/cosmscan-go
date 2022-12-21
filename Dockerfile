FROM golang:1.19.2-alpine as build_base
WORKDIR /src
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git gcc librdkafka-dev libc-dev
RUN cd /src && go mod download

FROM build_base AS builder

COPY . /src

ADD . /src
RUN cd /src && go build --tags musl -o /src/bin/ ./cmd/...

FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache jq

COPY --from=builder /src/bin /usr/local/bin

EXPOSE 8080
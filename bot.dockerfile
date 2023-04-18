FROM golang:1.20-alpine

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make && apk add bash

COPY bot/ bot/

COPY .env go.mod go.sum Makefile /app/

ENV GOBIN /go/bin

RUN go build -o main ./bot/cmd/main.go

CMD["./main"]
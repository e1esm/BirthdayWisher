FROM golang:1.20-alpine

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make && apk add bash

COPY bot/ bot/

COPY .env bot.dockerfile /app/

ENV GOBIN /go/bin

WORKDIR /app/bot

RUN go mod download && go mod tidy

RUN go build -o main ./cmd/main.go

CMD ["./main"]
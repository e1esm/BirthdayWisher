FROM golang:1.20-alpine

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make && apk add bash

COPY CongratulationsGenerator/ CongratulationsGenerator/
COPY congr_proto/ congr_proto/
COPY .env congratulator.dockerfile /app/

ENV GOBIN /go/bin

WORKDIR /app/CongratulationsGenerator

RUN go mod download && go mod tidy

RUN go build -o main ./cmd/main.go

CMD ["./main"]
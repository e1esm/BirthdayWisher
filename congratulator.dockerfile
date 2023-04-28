FROM golang:1.20-alpine

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make && apk add bash

COPY congratulationsGenerator/ CongratulationsGenerator/
COPY protobuf/bridge_to_API/gen_proto/ protobuf/bridge_to_API/gen_proto/
COPY .env congratulator.dockerfile /app/


ENV TZ="Europe/Moscow"
ENV GOBIN /go/bin

WORKDIR /app/CongratulationsGenerator

RUN go mod download && go mod tidy

RUN go build -o main ./cmd/main.go

CMD ["./main"]
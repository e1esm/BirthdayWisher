FROM golang:1.20-alpine

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make && apk add bash

COPY bridgeServer/ bridgeServer/
COPY protobuf/bridge_to_API/gen_proto/ protobuf/bridge_to_API/gen_proto/
COPY protobuf/bot_to_server/gen_proto/ protobuf/bot_to_server/gen_proto/
COPY .env bridge.dockerfile /app/

ENV TZ="Europe/Moscow"
ENV GOBIN /go/bin

WORKDIR /app/bridgeServer

RUN go mod download && go mod tidy

RUN go build -o main ./cmd/main.go

CMD ["./main"]
FROM golang:1.20-alpine

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make && apk add bash

COPY pdfGenerator/ pdfGenerator/
COPY protobuf/bridge_to_PDF-Generator/gen_proto/ protobuf/bridge_to_PDF-Generator/gen_proto/
COPY .env pdfGenerator.dockerfile /app/


ENV TZ="Europe/Moscow"
ENV GOBIN /go/bin

WORKDIR /app/pdfGenerator

RUN go mod download && go mod tidy

RUN go build -o main ./cmd/main.go

CMD ["./main"]
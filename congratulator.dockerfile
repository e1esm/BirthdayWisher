FROM joeka36/gorpc

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make && apk add bash && apk add --no-cache make protobuf-dev

COPY CongratulationsGenerator/ CongratulationsGenerator/

ENV GOBIN /go/bin

RUN go mod download && go mod tidy

RUN go build -o main ./cmd/main.go

CMD ["./main"]
FROM surnet/alpine-wkhtmltopdf:3.8-0.12.5-full as builder

FROM golang:1.20-buster
WORKDIR /app

RUN apt-get update && apt-get install libc-dev && apt-get install gcc && apt-get install make && apt-get install bash && apt-get install curl
RUN set -e; \
    apt-get update; \
    apt-get -y install cronjob; \
    apt-get install -y --no-install-recommends \
        apt-utils \
        ghostscript \
        fontforge \
        cabextract \
        wget \
        libjpeg62-turbo \
        xfonts-75dpi \
        xfonts-base; \
    wget https://gist.github.com/maxwelleite/10774746/raw/ttf-vista-fonts-installer.sh -q -O - | bash; \
    wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_amd64.deb ; \
    dpkg -i wkhtmltox_0.12.6-1.buster_amd64.deb; \
    rm  /app/wkhtmltox_0.12.6-1.buster_amd64.deb

COPY pdfGenerator/ pdfGenerator/
COPY protobuf/bridge_to_PDF-Generator/gen_proto/ protobuf/bridge_to_PDF-Generator/gen_proto/
COPY .env pdfGenerator.dockerfile /app/

COPY scripts/ scripts/

COPY --from=builder /bin/wkhtmltopdf /bin/wkhtmltopdf
COPY --from=builder /bin/wkhtmltoimage /bin/wkhtmltoimage

ENV TZ="Europe/Moscow"
ENV GOBIN /go/bin

WORKDIR /app/pdfGenerator

COPY cronjob /etc/cron.d/cronjob
RUN chmod 0644 /etc/cronjob.d/cronjob
RUN crontab /etc/cronjob.d/cronjob


RUN go mod download && go mod tidy
RUN go build -o main ./cmd/main.go

CMD cron && ./main
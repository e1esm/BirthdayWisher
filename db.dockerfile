FROM postgres:13.2-alpine

COPY /migrations/*.sql /docker-entrypoint-initdb.d/
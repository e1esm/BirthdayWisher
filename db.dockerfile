FROM postgres:13.2-alpine


ENV TZ="Europe/Moscow"
COPY /migrations/*.sql /docker-entrypoint-initdb.d/
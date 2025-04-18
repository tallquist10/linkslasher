ARG GO_VERSION=1.23

FROM golang:${GO_VERSION} as build
WORKDIR /app
COPY . /app
RUN ./build.sh $GO_VERSION

FROM alpine:latest AS db
WORKDIR /app
COPY sql/init.sql init.sql
RUN apk update && apk add sqlite
RUN sqlite3 linkslasher.db ".read init.sql"

FROM ubuntu:24.04
EXPOSE 8080
WORKDIR /app
COPY --from=build /app/linkslasher /app/linkslasher
COPY --from=build /app/templates/ /app/templates/
COPY --from=db /app/linkslasher.db /app/linkslasher.db
CMD ["./linkslasher"]
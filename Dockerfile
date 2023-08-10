FROM golang:1.20.6-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o task main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz
# Run stage
FROM alpine:3.13
WORKDIR /app
RUN apk add --no-cache bash
COPY --from=builder /app/task .
COPY --from=builder /app/migrate ./migrate
COPY migrate.sh .
COPY db/migration ./migration
RUN ["chmod", "+x", "/app/migrate.sh"]

EXPOSE 8080
CMD [ "/app/task", "serve" ]
ENTRYPOINT [ "/app/migrate.sh" ]
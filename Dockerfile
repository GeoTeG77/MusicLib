FROM golang:1.22-alpine as builder

RUN apk add --no-cache git build-base

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o musiclib cmd/main.go

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/musiclib /usr/local/bin/musiclib
COPY --from=builder /app/.env /app/.env 
COPY --from=builder /app/logs/ /app/logs/ 
COPY --from=builder /app/internal/migrations /app/internal/migrations

RUN chmod 644 /app/.env

EXPOSE 8080

CMD ["musiclib"]

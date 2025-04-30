FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /itbookworm ./cmd/api

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /itbookworm /app/itbookworm
COPY --from=builder /app/migrations /app/migrations

EXPOSE 4000
ENV DB_URL="postgres://user:password@db:5432/itbookworm?sslmode=disable" \
    APP_ENV="production"

CMD ["/app/itbookworm"]
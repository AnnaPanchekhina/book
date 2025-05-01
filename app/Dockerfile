FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /itbookworm ./cmd/api

FROM alpine:3.18

WORKDIR /app

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /itbookworm .
COPY --from=builder /app/migrations ./migrations
EXPOSE 4000
CMD ["/app/itbookworm"]
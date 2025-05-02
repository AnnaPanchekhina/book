FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /book ./cmd/api

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /book .
COPY --from=builder /app/migrations ./migrations
EXPOSE 4000
CMD ["./book"]
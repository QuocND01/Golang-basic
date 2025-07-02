# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY consumer/go.mod consumer/go.sum ./
COPY modules/model ../modules/model
RUN go mod download

COPY consumer/ ./
RUN go build -o app .

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]

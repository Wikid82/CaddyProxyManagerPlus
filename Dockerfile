# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o caddyproxymanager ./cmd/caddyproxymanager

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/caddyproxymanager .
COPY --from=builder /app/web ./web

# Create data directory
RUN mkdir -p /data /config

# Expose ports
EXPOSE 8080 80 443

# Set environment variables
ENV DATA_PATH=/data
ENV SERVER_PORT=8080
ENV CADDY_ADMIN_URL=http://localhost:2019

# Run the application
CMD ["./caddyproxymanager"]

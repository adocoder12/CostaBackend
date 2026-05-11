# Stage 1: Build
FROM golang:1.26-alpine AS builder



WORKDIR /app

# Copy dependency files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN mkdir -p /app/bin && \
    for cmd in cmd/*/; do \
        if [ -d "$cmd" ]; then \
            binary=$(basename $cmd); \
            echo "Building $binary..."; \
            CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$binary ./$cmd; \
        fi \
    done

# Stage 2: Final Runtime
FROM alpine:3.20


# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates


WORKDIR /app/

# Copy binary and migrations from builder
COPY --from=builder /app/bin/* ./

# Use the .env file in the container if provided by docker-compose
CMD ["./api"]
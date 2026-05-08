# Stage 1: Build
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copy dependency files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# Note: Changed ./cmd/api to ./cmd/main.go to match your previous folder structure
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o /app/bin/app \
    ./cmd/api/main.go
# Stage 2: Final Runtime
FROM alpine:3.20

# Create a non-privileged user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Install runtime dependencies (SSL certs for Supabase/AWS and Timezone data)
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy binary and migrations from builder
COPY --from=builder /app/bin/app .
# Ensure the path matches your actual migrations folder
COPY --from=builder /app/internal/repository/migrations ./migrations

# Change ownership to the non-root user
RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

# Use the .env file in the container if provided by docker-compose
CMD ["./app"]
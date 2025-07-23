# Use the official Go image as build environment
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server server.go

# Use alpine for the final image
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create app user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/server .

# Copy web assets
COPY --from=builder /app/web ./web

# Change ownership to app user
RUN chown -R appuser:appgroup /root/
USER appuser

# Expose port
EXPOSE 8343

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8343/ || exit 1

# Run the application
CMD ["./server"]

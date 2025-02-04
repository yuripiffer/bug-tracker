# Build from official Go Alpine image
FROM golang:1.21-alpine

# Install required system packages
RUN apk add --no-cache git

# Install go-junit-report globally
RUN go install github.com/jstemmer/go-junit-report/v2@latest

# Create non-root user
RUN adduser -D -u 1000 appuser && \
    mkdir -p /app && \
    chown -R appuser:appuser /app

# Configure environment variables
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

# Set working directory
WORKDIR /app

# Switch to non-root user
USER appuser
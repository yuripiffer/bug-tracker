# Build from official Go Alpine image
FROM golang:1.21-alpine

# Install required system packages
RUN apk add --no-cache git

# Install go-junit-report globally
RUN go install github.com/jstemmer/go-junit-report/v2@latest

# Create non-root user and cache directories
RUN adduser -D -u 1000 appuser && \
    mkdir -p /home/appuser/go/pkg/mod && \
    mkdir -p /home/appuser/.cache/go-build && \
    mkdir -p /go/pkg/mod && \
    chown -R appuser:appuser /home/appuser && \
    chown -R appuser:appuser /go/pkg/mod && \
    chmod -R 777 /go/pkg/mod

# Configure environment variables
ENV GOPATH=/home/appuser/go
ENV GOMODCACHE=/go/pkg/mod
ENV GOCACHE=/home/appuser/.cache/go-build
ENV PATH=$PATH:/home/appuser/go/bin

# Set working directory (separate from user home)
WORKDIR /app

# Switch to non-root user
USER appuser

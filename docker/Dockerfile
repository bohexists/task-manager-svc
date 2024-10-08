# Use official Golang image as the build stage
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to install dependencies
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o task-manager-svc ./cmd/main.go

# Use a minimal Alpine image for running the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=builder /app/task-manager-svc .

# Copy any necessary configuration files
COPY config/config.yaml ./config/

# Set environment variables (can be overridden in Docker Compose)
ENV DB_HOST=localhost
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASSWORD=secret
ENV DB_NAME=task_db

# Expose the gRPC port
EXPOSE 50051

# Run the compiled binary
CMD ["./task-manager-svc"]
# Use official Golang image as the build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to install dependencies
COPY ../go.mod ../go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY .. .

# Build the Go application
RUN go build -o task-manager-svc ./cmd/main.go

# Use a minimal Alpine image for running the application
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/task-manager-svc .

# Set environment variables
ENV DB_HOST=localhost
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASSWORD=secret
ENV DB_NAME=task_db

EXPOSE 50051

CMD ["./task-manager-svc"]
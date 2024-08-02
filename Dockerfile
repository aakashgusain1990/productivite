# Use the official Golang image to build the Go binary
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o main cmd/productivite/main.go

# Use a minimal base image to run the Go binary
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/config.yaml config.yaml

# Ensure the binary is executable
RUN chmod +x /root/main

# Debug: Check if the binary exists and is executable
RUN ls -l /root/main

# Expose the port that the Go server listens on (adjust if needed)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
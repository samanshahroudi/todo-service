# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Install git
RUN apk add --no-cache git

# Set the Go module proxy
ENV GOPROXY=https://proxy.golang.org,direct

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o todo-service ./cmd/main.go

# Stage 2: Minimal image for running the app
FROM alpine:latest AS runner

# Set environment variables (optional)
ENV GO_ENV=production

# Create a directory for the application
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/todo-service .

# Expose the port on which the application runs (if applicable)
EXPOSE 8080

# Command to run the application
CMD ["./todo-service"]
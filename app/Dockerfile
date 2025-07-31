# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to take advantage of Docker caching for dependencies
COPY go.mod go.sum ./

# Download Go module dependencies, only fetching if go.mod/go.sum change
RUN go mod download

# Copy the rest of the application source code
COPY *.go ./

# Build the Go application
# CGO_ENABLED=0 creates a statically linked binary (no C libraries needed at runtime)
# GOOS=linux ensures compatibility with Linux containers
# -ldflags="-w -s" reduces binary size by omitting debug info and symbol table
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# Stage 2: Create a lightweight image for the compiled binary
FROM alpine:latest

# Set working directory for the application
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application listens on (informational)
EXPOSE 3000

# Run the Go application when the container starts
CMD ["./main"]

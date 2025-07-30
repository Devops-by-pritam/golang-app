# Stage 1: Build the Go app
FROM golang:1.21 as builder

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY app/go.mod .
COPY app/go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY app/*.go .  # Make sure your .go files are inside the app/ folder

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Stage 2: Run the compiled binary in a lightweight image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the compiled binary from builder stage
COPY --from=builder /app/main .

# Expose port the app listens on
EXPOSE 3000

# Run the app
CMD ["./main"]

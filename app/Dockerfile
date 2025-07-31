# Stage 1: Build the Go application
# We use a specific version of Go on Alpine Linux to create the binary.
# This stage is called 'builder'.
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container for this stage.
WORKDIR /app

# Copy the go.mod and go.sum files first. This is a critical step for caching.
# Docker will only re-run 'go mod download' if these files change.
COPY go.mod go.sum ./

# Download Go module dependencies. This command is fast if the cache is hit.
RUN go mod download

# Copy the rest of the application source code into the working directory.
COPY *.go ./

# Build the Go application binary.
# CGO_ENABLED=0: Creates a statically linked binary without C libraries, for portability.
# GOOS=linux: Ensures the binary is compiled for a Linux environment.
# -ldflags="-w -s": Strips debug symbols to significantly reduce the binary size.
# -o main .: Names the output binary 'main'.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# -----------------------------------------------------------------------------

# Stage 2: Create a lightweight final image
# We use the smallest possible base image, a minimal alpine, to hold only the binary.
FROM alpine:latest

# Set the working directory for the final application.
WORKDIR /app

# Copy the compiled 'main' binary from the 'builder' stage into this new image.
# We use the '--from=builder' flag to reference the first stage.
COPY --from=builder /app/main .

# Expose the port on which the application will listen. This is informational.
EXPOSE 3000

# Set the command to run the application when the container starts.
CMD ["./main"]

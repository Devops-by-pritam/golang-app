FROM golang:1.21 as builder

WORKDIR /app
COPY . .

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 3000
CMD ["./main"]

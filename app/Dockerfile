FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN go mod init golang-app && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 3000
CMD ["./main"]

# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o /app-server main.go

# Stage 2: Create a minimal final image
FROM alpine:3.18

COPY --from=builder /app-server /usr/local/bin/app-server

EXPOSE 8080

CMD ["/usr/local/bin/app-server"]
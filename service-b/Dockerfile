# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go

# Final image
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/server .

EXPOSE 8081
CMD ["./server"]
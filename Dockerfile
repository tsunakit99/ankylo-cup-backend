FROM golang:1.22-alpine3.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# .envファイルをコピー
COPY .env .env

COPY . .
RUN go build -o server ./cmd/server/main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/.env .env
EXPOSE 50051
CMD ["./server"]

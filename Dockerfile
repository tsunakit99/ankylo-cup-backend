# builder stage
FROM golang:1.22-alpine3.21 AS builder
WORKDIR /app

# go.mod, go.sum のコピー・ダウンロード
COPY go.mod go.sum ./
RUN go mod download

# .env と credentials.json のコピー
COPY .env .env
COPY internal/auth/credentials.json internal/auth/credentials.json

# ソースコード全体をコピー
COPY . .

RUN go build -o server ./cmd/server/main.go

# final stage
FROM alpine:3.17

WORKDIR /app
# 実行ファイルと.envをコピー
COPY --from=builder /app/server .
COPY --from=builder /app/.env .env

# credentials.jsonをコピー
COPY --from=builder /app/internal/auth/credentials.json internal/auth/credentials.json

# ポート公開など
EXPOSE 50051

CMD ["./server"]

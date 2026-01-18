# =============================================================================
# CI: CI検証用 （make check を実行）
# =============================================================================
FROM golang:alpine AS ci

WORKDIR /app

RUN apk add --no-cache make
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install golang.org/x/tools/cmd/goimports@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["make", "check"]

# =============================================================================
# builder: 本番用マルチステージビルドの前段 （本番用バイナリのビルド）
# =============================================================================
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o bff .

# =============================================================================
# runner: 本番用マルチステージビルドの後段 （最小イメージでバイナリ実行）
# =============================================================================
FROM gcr.io/distroless/static-debian12 AS runner

WORKDIR /app
COPY --from=builder /app/bff .

EXPOSE 8000

CMD ["./bff"]

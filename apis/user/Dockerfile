# =============================================================================
# CI: CI検証用 （make check を実行）
# =============================================================================
FROM golang:1.25-alpine AS ci

WORKDIR /app

RUN apk add --no-cache make diffutils
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8
RUN go install golang.org/x/tools/cmd/goimports@v0.43.0
RUN go install github.com/ogen-go/ogen/cmd/ogen@v1.18.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.30.0
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.19.1

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["make", "check"]

# =============================================================================
# builder: 本番用マルチステージビルドの前段 （本番用バイナリのビルド）
# =============================================================================
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o user .

# =============================================================================
# runner: 本番用マルチステージビルドの後段 （最小イメージでバイナリ実行）
# =============================================================================
FROM gcr.io/distroless/static-debian12 AS runner

WORKDIR /app
COPY --from=builder /app/user .

EXPOSE 8000

CMD ["./user"]

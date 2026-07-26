FROM mcr.microsoft.com/devcontainers/go:1.25

RUN go install github.com/ogen-go/ogen/cmd/ogen@v1.18.0 \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8 \
    && go install github.com/air-verse/air@v1.64.0 \
    && go install golang.org/x/tools/cmd/goimports@v0.43.0 \
    && go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.30.0 \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.19.1 \
    && go install go.uber.org/mock/mockgen@v0.6.0 \
    && chmod -R g+w /go/pkg/

# ベースイメージ同梱の Yarn APT source が古い GPG キーで失敗するため削除
# TODO: ベースイメージを golang:1.26-bookworm に移行して根治
RUN rm -f /etc/apt/sources.list.d/yarn.list /etc/apt/sources.list.d/yarn.sources \
    && apt-get update && apt-get install -y \
    curl \
    git \
    wget \
    jq \
    ripgrep \
    tree \
    make \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

FROM mcr.microsoft.com/devcontainers/go:1.21

# 追加のGoツールをインストール（必要に応じて）
# RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

RUN apt-get update && apt-get install -y \
    curl \
    git \
    wget \
    jq \
    ripgrep \
    tree \
    make \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

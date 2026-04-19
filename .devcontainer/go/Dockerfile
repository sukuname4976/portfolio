FROM mcr.microsoft.com/devcontainers/go:1.25

RUN go install github.com/ogen-go/ogen/cmd/ogen@v1.18.0 \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8 \
    && go install github.com/air-verse/air@v1.64.0 \
    && go install golang.org/x/tools/cmd/goimports@v0.43.0 \
    && chmod -R g+w /go/pkg/

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

# Terraform Development Environment Dockerfile
FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Terraform、AWS CLI、Azure CLIはfeaturesでインストールされるため、
# 追加のツールのみインストール
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    jq \
    unzip \
    && rm -rf /var/lib/apt/lists/*


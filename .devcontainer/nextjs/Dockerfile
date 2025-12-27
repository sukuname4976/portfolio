FROM mcr.microsoft.com/devcontainers/typescript-node:18

RUN apt-get update && apt-get install -y \
    curl \
    git \
    wget \
    jq \
    ripgrep \
    tree \
    make \
    && rm -rf /var/lib/apt/lists/*

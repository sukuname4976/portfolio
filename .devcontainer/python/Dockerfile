FROM mcr.microsoft.com/devcontainers/python:3.11

RUN curl -LsSf https://astral.sh/uv/install.sh | sh

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

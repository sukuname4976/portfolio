# Python Development Environment Dockerfile
FROM mcr.microsoft.com/devcontainers/python:3.11

# 追加のPythonツールをインストール（必要に応じて）
# RUN pip install --no-cache-dir \
#     pytest \
#     black \
#     ruff

# その他の開発ツール
RUN apt-get update && apt-get install -y \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*


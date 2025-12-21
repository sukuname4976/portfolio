# Node.js/Next.js Development Environment Dockerfile
FROM mcr.microsoft.com/devcontainers/typescript-node:18

# グローバルパッケージのインストール（必要に応じて）
# RUN npm install -g \
#     typescript \
#     ts-node

# その他の開発ツール
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    && rm -rf /var/lib/apt/lists/*


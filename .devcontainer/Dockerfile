# ルートレベルの共通開発環境用Dockerfile
# 必要に応じてカスタマイズ

FROM mcr.microsoft.com/devcontainers/base:ubuntu

# 共通ツールのインストール
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Nxはnpm経由でインストールされるため、Node.jsのインストールで対応
# package.jsonのdevDependenciesに含まれている

# Terraformのインストール
RUN wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | tee /usr/share/keyrings/hashicorp-archive-keyring.gpg \
    && echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list \
    && apt-get update && apt-get install -y terraform \
    && rm -rf /var/lib/apt/lists/*

# Node.jsのインストール（Nxに必要）
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs \
    && rm -rf /var/lib/apt/lists/*

# Pythonのインストール（必要に応じて）
# RUN apt-get update && apt-get install -y python3 python3-pip \
#     && rm -rf /var/lib/apt/lists/*


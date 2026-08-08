# :ubuntu(最新追従)は 26.04 LTS(resolute)へ跳ね、apt-key 廃止で
# google-cloud-cli feature の鍵登録が失敗する。最新 LTS を codename で固定する。
FROM mcr.microsoft.com/devcontainers/base:resolute

# gcloud は feature(apt-key 依存で 26.04 で動かない)をやめ、
# apt-key 非依存の keyring 方式 (signed-by) で自前導入する。
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    ca-certificates \
    gnupg \
    git \
    wget \
    jq \
    ripgrep \
    tree \
    make \
    && curl -fsSL https://packages.cloud.google.com/apt/doc/apt-key.gpg \
        | gpg --dearmor -o /usr/share/keyrings/cloud.google.gpg \
    && echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" \
        > /etc/apt/sources.list.d/google-cloud-sdk.list \
    && apt-get update && apt-get install -y --no-install-recommends google-cloud-cli \
    && rm -rf /var/lib/apt/lists/*

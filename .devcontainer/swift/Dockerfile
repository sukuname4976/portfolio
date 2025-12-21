# Swift/iOS Development Environment Dockerfile
FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Swiftのインストール（Linux環境用）
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    build-essential \
    libicu-dev \
    && rm -rf /var/lib/apt/lists/*

# Swiftのインストール（公式リポジトリから）
RUN SWIFT_VERSION=5.9 && \
    wget -q -O - https://swift.org/keys/all-keys.asc | gpg --import - && \
    wget -q https://swift.org/builds/swift-${SWIFT_VERSION}-release/ubuntu2204/swift-${SWIFT_VERSION}-RELEASE/swift-${SWIFT_VERSION}-RELEASE-ubuntu22.04.tar.gz && \
    tar xzf swift-${SWIFT_VERSION}-RELEASE-ubuntu22.04.tar.gz && \
    mv swift-${SWIFT_VERSION}-RELEASE-ubuntu22.04 /opt/swift && \
    rm swift-${SWIFT_VERSION}-RELEASE-ubuntu22.04.tar.gz

ENV PATH="/opt/swift/usr/bin:${PATH}"

# 注意: 完全なiOS開発にはmacOSとXcodeが必要です
# このDockerfileはLinux環境での基本的なSwift開発をサポートします


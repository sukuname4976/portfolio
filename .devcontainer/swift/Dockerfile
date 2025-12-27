FROM mcr.microsoft.com/devcontainers/base:ubuntu

RUN apt-get update && apt-get install -y \
    wget \
    build-essential \
    libicu-dev \
    && rm -rf /var/lib/apt/lists/*

RUN SWIFT_VERSION=5.9 && \
    wget -q -O - https://swift.org/keys/all-keys.asc | gpg --import - && \
    wget -q https://swift.org/builds/swift-${SWIFT_VERSION}-release/ubuntu2204/swift-${SWIFT_VERSION}-RELEASE/swift-${SWIFT_VERSION}-RELEASE-ubuntu22.04.tar.gz && \
    tar xzf swift-${SWIFT_VERSION}-RELEASE-ubuntu22.04.tar.gz && \
    mv swift-${SWIFT_VERSION}-RELEASE-ubuntu22.04 /opt/swift && \
    rm swift-${SWIFT_VERSION}-RELEASE-ubuntu22.04.tar.gz

ENV PATH="/opt/swift/usr/bin:${PATH}"

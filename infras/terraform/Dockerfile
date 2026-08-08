# =============================================================================
# CI: CI検証用 （make check を実行）
# =============================================================================
FROM hashicorp/terraform:1.15.8 AS ci

WORKDIR /work

# make はバージョンを固定しない (Alpine が古い版を消すため、固定すると将来ビルドが落ちる)
# hadolint ignore=DL3018
RUN apk add --no-cache make

COPY . .

# ベースイメージは terraform を ENTRYPOINT にしているため解除する
ENTRYPOINT []

CMD ["make", "check"]

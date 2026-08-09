# AGENTS.md

## 概要

ポートフォリオの monorepo。nx で管理し、devcontainer 上で開発する。

- `apis/` - バックエンド API (Go, Python)
- `clients/` - フロントエンド (Web)
- `infras/` - インフラ (Terraform)
- `libraries/` - 共有ライブラリ

## 作業の進め方

- このファイルを読み込んだら、応答の冒頭に次の 1 行を出力する
  - `ドキュメントの指示に従い repository-overview/agents-base.md を参照したことを通知します`
- 着手前に `find-adr` でその領域の決定を確認する
- 選択肢を比べて片方を捨てた判断は `record-adr` で ADR に残す
- 原則として既存のコードスタイルに従う

# Terraform

## 概要

インフラストラクチャ定義。
現時点では GitHub リポジトリの設定を IaC として管理する。

## 技術スタック

- Terraform 1.15
- integrations/github : GitHub provider

## ディレクトリ構成

```text
infras/terraform/
├── main.tf                # provider 定義·backend 方針
├── variables.tf           # 変数定義
├── github.tf              # GitHub リポジトリの設定
├── outputs.tf             # 出力定義
├── load-env.sh            # 認証情報を環境変数へ読み込む
├── .terraform.lock.hcl    # provider バージョンの固定
├── Dockerfile             # CI 検証用コンテナ定義
├── docker-compose.ci.yaml # CI 用 Compose
├── Makefile               # コマンド
└── project.json           # Nx プロジェクト設定
```

## 開発環境のセットアップ

1. VSCode で `.devcontainer/terraform` の devcontainer を開く
2. ターミナルで以下を実行

```bash
cd infras/terraform

# GitHub 認証 (コンテナに入り直すたびに必要)
gh auth login

# 環境変数の読み込み (source して現在のシェルに反映させる)
source ./load-env.sh

# 初期化
terraform init
```

## make コマンド

- `make init` : 初期化 (provider の取得)
- `make plan` : 差分確認
- `make apply` : 適用 (変更をコミットしてから実行する)
- `make fmt` : フォーマット
- `make check` : CI チェック (フォーマット確認 + 検証)
- `make prepare` : コミット前の品質担保処理 (フォーマット + 検証)

`make check` は GitHub API を叩かないため認証情報を必要としない。
CI では `docker-compose.ci.yaml` 経由で同じ `make check` を実行するので、
ローカルと CI で検証内容が一致する。

## 認証

- GitHub provider は環境変数 `GITHUB_TOKEN` を参照する
- `load-env.sh` はトークンの値ではなく取得コマンド (`gh auth token`) を
  記述しているため、秘密情報がファイルにもディスクにも残らない
- devcontainer はホストと home を共有しないため、コンテナに入り直すたびに
  `gh auth login` が必要になる

## 運用ルール

### コミットしてから apply する

変更は必ずコミットしてから `terraform apply` する。

1. `terraform plan` で差分を確認する
2. 変更をコミットする
3. `terraform apply` で適用する

理由:

- IaC はコードを環境の唯一の正とする仕組みである
- 未コミットのまま apply すると、環境にだけ存在してコードに残らない変更が
  生まれ、コードと実環境が乖離する
- 乖離すると現在の環境が何によって作られたのかを追えなくなり、
  IaC が成立しなくなる

### apply は手動ローカル実行のみ

- state はローカル管理とし、CI では実行しない
- リモート化する場合は `main.tf` に backend を定義する

## 管理対象

### main ブランチの保護

`github.tf` の `github_repository_ruleset` で管理する。

- 直 push を禁止し、変更を PR 経由に限定する
- force-push を禁止する
- ブランチの削除を禁止する

`bypass_actors` は定義していないため、管理者も直 push できない。

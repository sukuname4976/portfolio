# Terraform

## 概要

インフラストラクチャ定義。Terraform でコードとして管理する。

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
- devcontainer はホストと home を共有しないため、コンテナに入り直すたびに
  `gh auth login` が必要になる

秘密をファイルに残さず取得コマンドだけを書く方針は `infra-terraform-0002` を
参照。

## 適用の手順

1. `terraform plan` で差分を確認する
2. 変更をコミットする
3. `terraform apply` で適用する

`apply` と `destroy` は人間が実行する。エージェントが進めてよいのは `plan`
まで。

この手順を定めた理由と `permissions.deny` による担保は
`infra-terraform-0001` を参照。

## 管理対象

- main ブランチの保護 (`github.tf`)。規則の内容と選定の理由は
  `infra-terraform-0003` を参照

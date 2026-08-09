---
id: infra-terraform-0004
summary: |-
  バージョン制約だけでは環境ごとに異なる provider が入りうる。.terraform.lock.hcl を
  コミットしてバージョンを固定する。ロックファイルは init したプラットフォームの
  ハッシュしか持たないため、開発環境と CI の両方を記録し、CI では readonly で
  古さを検出する。
status: 採用
level: choice
created: 2026-08-09
updated: 2026-08-09
version: 1.0.0
---

# infra-terraform-0004. .terraform.lock.hcl をコミットする

## 決定事項

- `.terraform.lock.hcl` を追跡の対象とし、コミットする
- `terraform providers lock` で、開発環境と CI の両方のプラットフォームの
  ハッシュを記録する
- CI では `terraform init -lockfile=readonly` を使い、ロックファイルが
  古い場合は失敗させる

## コンテキスト

`infras/terraform/.gitignore` は当初 `.terraform.lock.hcl` を除外していた。
Terraform の公式はロックファイルをコミットすることを推奨しており、`init` の
実行時にも同じ趣旨の案内が出る。

ロックファイルは `init` を実行したプラットフォームの `h1:` ハッシュしか
記録しない。この開発環境は devcontainer 上の linux_arm64 で、GitHub Actions の
ランナーは linux_amd64 である。

provider のバージョン制約は `main.tf` に `~> 6.0` として書かれている。

## 却下事項

- **gitignore したままにする**: ロックファイルを管理せずに済み、プラット
  フォームの違いも気にしなくてよいが、`~> 6.0` のような制約だけでは環境ごとに
  異なるバージョンが入りうる。`fmt` の結果や検証の挙動が環境で変わる
- **開発環境のプラットフォームだけ記録する**: `terraform init` を実行するだけ
  で済むが、CI のプラットフォームのハッシュが無い。`-lockfile=readonly` を
  使うとロックファイルを更新できずに失敗する
- **CI で `-lockfile=readonly` を使わない**: ロックファイルが古くても CI が
  通るが、更新の漏れに気づけない。コミットして固定する意味が薄れる

## 前提と見直し条件

- 開発環境が linux_arm64、CI が linux_amd64 であることを前提とする。
  どちらかが変わったとき、記録するプラットフォームを見直す
- provider を registry から取得することを前提とする。ミラーや private registry
  を使うようになったとき、ハッシュの扱いを見直す

## 関連資料

- [[infra-terraform-0001-policy-terraform-with-manual-apply]] :
  インフラを Terraform で管理する方針
- `infras/terraform/.terraform.lock.hcl` : 実体
- Issue #9 : 本決定を行った作業

## 変更履歴

- v1.0.0 (2026-08-09): 初版

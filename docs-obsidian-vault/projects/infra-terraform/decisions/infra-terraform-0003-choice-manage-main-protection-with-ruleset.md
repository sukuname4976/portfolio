---
id: infra-terraform-0003
summary: |-
  main が無保護で直 push が通り、過去に feature ブランチの push が main へ直行する
  事故が起きていた。保護は従来の branch protection ではなく後継の ruleset で管理し、
  PR を必須にして force-push とブランチ削除を禁止する。承認必須数は 0、
  bypass_actors は定義しない。
status: 採用
level: choice
created: 2026-08-09
updated: 2026-08-09
version: 1.0.0
---

# infra-terraform-0003. main ブランチ保護を ruleset で管理する

## 決定事項

- main ブランチの保護は `github_repository_ruleset` で管理する
- 対象は `~DEFAULT_BRANCH` で指定し、ブランチ名を直接書かない
- 次の規則を適用する
  - `pull_request` を必須とし、承認必須数は 0 とする
  - `non_fast_forward` により force-push を禁止する
  - `deletion` によりブランチの削除を禁止する
- `bypass_actors` は定義しない。管理者も直接 push できない状態を保つ

## コンテキスト

main ブランチは無保護で、直接 push が通る状態だった。過去に upstream の設定が
`origin/main` になっていたことで、feature ブランチのつもりの push が main へ
直行する事故が起きている。

GitHub にはブランチを保護する仕組みが 2 つある。従来の branch protection と、
後継として位置づけられている ruleset である。ruleset は `~DEFAULT_BRANCH` の
ような組み込みのパターンで対象を指定できる。

このリポジトリはソロで運用しており、GitHub の仕様上、自分が作成した PR を
自分で承認することはできない。

## 却下事項

- **従来の branch protection で管理する**: 情報が多く provider の対応も枯れて
  いるが、GitHub が ruleset を後継として位置づけている。対象の指定も
  ブランチ名の直書きになり、デフォルトブランチを変えたときに追従しない
- **承認必須数を 1 以上にする**: レビューを経ない変更を止められるが、自分の
  PR を自分で承認できないため、ソロ運用では自分の PR がマージできなくなる
- **`bypass_actors` に管理者を入れる**: 緊急時に直接 push できるが、事故を
  止める最後の砦が失われる。過去に起きた事故はまさに直接 push だった
- **ブランチ名を直接指定する**: 対象が一目で分かるが、デフォルトブランチを
  変更したときに設定が追従せず、保護が外れたことに気づけない

## 前提と見直し条件

- ソロ運用であることを前提とする。自分以外のコミッターが加わったとき、
  承認必須数を見直す
- 緊急の直接 push が必要になったことはない。必要になったとき、
  `bypass_actors` を追加するかを検討する
- GitHub が ruleset を後継として維持していることを前提とする。位置づけが
  変わったとき、管理の方法を見直す

## 関連資料

- [[infra-terraform-0001-policy-terraform-with-manual-apply]] :
  インフラを Terraform で管理する方針
- `infras/terraform/github.tf` : 実装
- Issue #9 : 本決定を行った作業

## 変更履歴

- v1.0.0 (2026-08-09): 初版

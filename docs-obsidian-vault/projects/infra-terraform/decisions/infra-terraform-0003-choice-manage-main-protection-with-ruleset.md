---
id: infra-terraform-0003
summary: |-
  main が無保護で直 push が通り、過去に feature ブランチの push が main へ直行する
  事故が起きていた。保護は従来の branch protection ではなく後継の ruleset で管理し、
  PR を必須にして force-push とブランチ削除を禁止する。CI の 2 ジョブと
  AI レビューの 1 ジョブの成功を必須のステータスチェックとする。承認必須数は 0、
  bypass_actors は定義しない。
status: 採用
level: choice
created: 2026-08-09
updated: 2026-09-21
version: 1.1.0
---

# infra-terraform-0003. main ブランチ保護を ruleset で管理する

## 決定事項

- main ブランチの保護は `github_repository_ruleset` で管理する
- 対象は `~DEFAULT_BRANCH` で指定し、ブランチ名を直接書かない
- 次の規則を適用する
  - `pull_request` を必須とし、承認必須数は 0 とする
  - `non_fast_forward` により force-push を禁止する
  - `deletion` によりブランチの削除を禁止する
  - `required_status_checks` により、CI の 2 ジョブと AI レビューの 1 ジョブの
    成功を必須のステータスチェックとする。対象は
    `Nx 変更検出したプロジェクトへ CI を実行`、
    `リポジトリ全体へ Markdown Lint を実行`、
    `Claude Code による PR コードレビューを実行` で、`context` はワークフローの
    ジョブ名と一致させ、`integration_id` は指定しない
  - `strict_required_status_checks_policy` は false とする
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

初版の ruleset は PR を必須にするだけで、CI の成功をマージの条件にして
いなかった。CI が落ちていても PR をマージできる状態で、Issue #36 で規範を
仕組みで強制する方針を書く途中で判明した。CI は `.github/workflows/ci.yaml` の
2 ジョブからなり、AI レビューは `.github/workflows/claude-code-PR-review.yaml` の
1 ジョブで PR ごとに実行される。GitHub Actions のチェックはジョブ名を
`context` として扱う。
`strict_required_status_checks_policy` を true にすると、PR のブランチが main の
最新を取り込んでいない限りマージできない。

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
- **必須のステータスチェックを設けない**: ruleset の設定が短く済み、CI の
  ジョブ名を変えても ruleset に追従させる手間がないが、CI が落ちていても PR を
  マージできる。マージ前の自動検証を人が見落とすと、壊れた変更が main に入る。
  AI レビューの指摘が返る前にマージすることも止められない
- **`strict_required_status_checks_policy` を true にする**: main の最新と
  合わせた状態で CI を通した変更だけがマージされ、個別には通るが合わせると
  壊れる変更を防げるが、main が進むたびに PR ごとに取り込み直して CI を
  やり直すことになる。ソロ運用では並行して進む PR が少なく、その手間が
  見合わない

## 前提と見直し条件

- ソロ運用であることを前提とする。自分以外のコミッターが加わったとき、
  承認必須数を見直す
- 緊急の直接 push が必要になったことはない。必要になったとき、
  `bypass_actors` を追加するかを検討する
- GitHub が ruleset を後継として維持していることを前提とする。位置づけが
  変わったとき、管理の方法を見直す
- 必須のステータスチェックはワークフローのジョブ名で指定している。ジョブ名を
  変えたとき、またはジョブを増減したとき、`required_check` を合わせて変える
- 並行して進む PR が少ないことを前提とする。個別には CI が通るのに main へ
  マージすると壊れる変更が入ったとき、
  `strict_required_status_checks_policy` を true にするかを見直す

## 関連資料

- [[infra-terraform-0001-policy-terraform-with-manual-apply]] :
  インフラを Terraform で管理する方針
- [[repository-overview-0005-design-automated-pr-review]] :
  すべての PR に AI レビューを通す決定。本決定はその成功をマージの条件にする
- `infras/terraform/github.tf` : 実装
- `.github/workflows/ci.yaml` : 必須のステータスチェックとする CI のジョブ
- `.github/workflows/claude-code-PR-review.yaml` : 必須のステータスチェックと
  する AI レビューのジョブ
- Issue #9 : 本決定を行った作業
- Issue #41 : 必須のステータスチェックを追加した作業

## 変更履歴

- v1.1.0 (2026-09-21): CI の 2 ジョブと AI レビューの 1 ジョブの成功を必須の
  ステータスチェックとする規則を追加した
- v1.0.0 (2026-08-09): 初版

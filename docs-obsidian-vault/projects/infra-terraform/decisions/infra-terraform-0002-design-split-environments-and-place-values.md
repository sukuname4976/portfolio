---
id: infra-terraform-0002
summary: |-
  root module が 1 つで state もローカルに 1 つのため、dev と prod を表現できない。
  環境ごとにディレクトリを分けて state を持たせ、環境固有の値は tfvars に置いて
  コミットする。秘密だけは値を書かず、取得コマンドを書いてファイルに残さない。
status: 採用
level: design
created: 2026-08-09
updated: 2026-08-09
version: 1.0.0
---

# infra-terraform-0002. dev と prod を env ディレクトリで分け値と秘密の置き場所を定める

## 決定事項

### 環境の分け方

- dev と prod は `envs/<環境>/` のディレクトリで分け、環境ごとに backend と
  state を持たせる
- 共通して使う部分は `modules/` に置き、各環境から呼ぶ
- 環境に属さない設定は `global/` に置く

### 値の置き場所

- 環境ごとの値は `envs/<環境>/terraform.tfvars` に書き、コミットする
- 適用対象を決める値 (プロジェクト ID、オーナー、リポジトリ名など) には
  `default` を持たせない
- 調整のための値 (リージョン、インスタンス数など) には `default` を持たせてよい

### 秘密の置き場所

- 秘密はファイルに保存しない。値ではなく取得するコマンドを書く
- 取得コマンドは `load-env.sh` に置き、`source` して環境変数へ流す
- `.gitignore` の対象は秘密を含みうる上書き (`*.auto.tfvars` など) に限り、
  環境定義の `terraform.tfvars` は除外しない

## コンテキスト

`infras/terraform` は root module が 1 つで、state もローカルに 1 つしかない。
この形では dev と prod を同じ state に入れることになり、`plan` が常に両方を
対象にする。Issue #33 で Vercel、Supabase、Cloud Run を dev と prod の両方に
作る予定があり、現状では表現できない。

`infras/terraform/.gitignore` は `*.tfvars` を丸ごと除外している。
`variables.tf` の `github_owner` と `repository_name` は `default` を持っており、
値を指定しなくても適用できる状態にある。

認証は `load-env.sh` が `export GITHUB_TOKEN="$(gh auth token)"` を実行する形で
実装済みで、トークンの値はファイルにもディスクにも残らない。`gh` が発行する
トークンは短命で自動更新されるため、有効期限の管理も発生しない。

## 却下事項

### 環境の分け方

- **`terraform workspace` で分ける**: 設定が 1 つで済みディレクトリも増えないが、
  選択中の workspace が画面に出ない。適用は人間が手で行うため、対象の環境が
  パスに現れないと確認の手がかりを失い、prod を選んだまま適用する事故が起きる

### 値の置き場所

- **tfvars を gitignore したままにする**: 秘密が混ざっても漏れないが、環境の
  定義がリポジトリの外に出る。prod が何によって作られたかがコードに残らず、
  クローンした直後は `plan` すら通らない
- **適用対象を `variables.tf` の `default` に固定する**: 追加の設定なしで
  `plan` が通るが、値を指定しないと既定の対象へ適用される。環境が増えたとき
  「指定しなければ prod になる」形になる
- **`default` をすべて外す**: 対象を必ず明示させられるが、調整のための値まで
  毎回書くことになる。`validate` は通るが `plan` が止まり、確認の手間が増える

### 秘密の置き場所

- **`.tfvars` に値を書いて gitignore する**: 設定が 1 箇所にまとまるが、秘密が
  ディスク上に平文で残る。トークンの有効期限も自分で管理することになる
- **環境変数へ値を直接書く**: 取得の手間がないが、シェルの履歴や設定ファイルに
  残り、消したつもりでも残留する

## 前提と見直し条件

- dev と prod の 2 環境を想定している。3 つ以上に増えたとき、ディレクトリの
  構成を見直す
- state をローカル管理していることを前提とする。prod を扱うようになったとき、
  リモート state への移行を検討する
- `gh auth token` が短命なトークンを返すことを前提とする。取得の方法が
  変わったとき、`load-env.sh` を見直す
- 秘密を扱う provider が GitHub だけであることを前提とする。GCP や Supabase を
  追加したとき、同じ方式で取得コマンドを書けるかを確認する

## 関連資料

- [[infra-terraform-0001-policy-terraform-with-manual-apply]] :
  情報をリポジトリへ集める方針と、適用の手順を定める決定
- Issue #33 : 本決定を実装するタスク
- `infras/terraform/load-env.sh` : 秘密の取得方法の実装

## 変更履歴

- v1.0.0 (2026-08-09): 初版

#!/usr/bin/env bash
# Terraform 実行用の環境変数を読み込むスクリプト。
#
# 使い方（サブシェルではなく現在のシェルに反映させるため source する）:
#   source ./load-env.sh
#   terraform plan
#
# 設計方針:
#   トークンの「値」ではなく「取得コマンド」を書くことで、
#   秘密情報をこのファイルにもディスクにも残さない。
#   gh のトークン（gho_）は短命かつ自動更新されるため、有効期限を気にする必要もない。

#   （source される前提なので `set -e` 等は入れない。呼び出し元のシェル設定を汚さないため）

if ! command -v gh >/dev/null 2>&1; then
  echo 'エラー: gh コマンドが見つかりません。' >&2
  return 1 2>/dev/null || exit 1
fi

if ! gh auth status >/dev/null 2>&1; then
  echo 'エラー: gh が未認証です。`gh auth login` を実行してください。' >&2
  return 1 2>/dev/null || exit 1
fi

# GitHub provider は環境変数 GITHUB_TOKEN を自動で参照する。
GITHUB_TOKEN="$(gh auth token)"
export GITHUB_TOKEN

echo 'GITHUB_TOKEN を設定しました。'

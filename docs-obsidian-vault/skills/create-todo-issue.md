---
name: create-todo-issue
description: シンプルな todo Issue を作成する。フォーマット指定なし、gh issue create で作成。
---

# create-todo-issue スキル

シンプルな todo Issue を作成する。

## 手順

1. ユーザーから todo 内容をヒアリング
2. `gh issue create` で Issue を作成

## Issue 作成実行の全体の例

```bash
gh issue create --title "READMEにセットアップ手順を追記" --label "todo" --body "$(cat <<'EOF'
## 概要

README.md にローカル開発環境のセットアップ手順を追記する。

## 備考

特に無し
EOF
)"
```

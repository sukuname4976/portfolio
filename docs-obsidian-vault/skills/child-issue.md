---
name: child-issue
description: 子 Issue（Task）を作成する。親 Issue に紐づけてテンプレートに従って作成し、gh issue create で作成。
---

# child-issue スキル

子 Issue（Task）を作成する。

## 手順

1. 親 Issue の番号を確認
2. タスク内容をヒアリング
3. テンプレートに従って Issue 本文を作成
4. `gh issue create` で Issue を作成
5. 親 Issue に子 Issue へのリンクを追記

## 参照テンプレート

`.github/ISSUE_TEMPLATE/child_task.yaml`

## Issue 作成実行の全体の例

<!-- markdownlint-disable MD013 -->

```bash
gh issue create --title "[Child Task]: ログイン API を実装" --label "Child Task" --body "$(cat <<'EOF'
## 📋 概要

ユーザー認証機能のログイン API エンドポイントを実装する。

## 📦 対象プロジェクト

- [x] 🔀 apis/bff

## 🏷️ タスク種別

- [x] ⚙️ バックエンド開発

## ✅ 子タスク完了条件

- [ ] POST /api/auth/login が Swagger 上で正常に動作確認が完了している
- [ ] JWT ミドルウェアが実装され、手元での動作確認が完了している
- [ ] 単体テスト・モックが用意されている
- [ ] CI (test, lint, format) が全て正常に終了している

## 🔗 このタスクに紐付く親タスク

#10

## 💬 備考

特に無し
EOF
)"
```

<!-- markdownlint-enable MD013 -->

---
name: parent-issue
description: 親 Issue(Epic/Feature)を作成する。テンプレートに従って Issue 本文を作成し、gh issue create で作成。
---

# parent-issue スキル

親 Issue（Epic/Feature）を作成する。

## 手順

1. ユーザーから要件をヒアリング
2. テンプレートに従って Issue 本文を作成
3. `gh issue create` で Issue を作成

## 参照テンプレート

`.github/ISSUE_TEMPLATE/parent_task.yaml`

## Issue 作成実行の全体の例

<!-- markdownlint-disable MD013 -->

```bash
gh issue create --title "[Parent Task]: ユーザー認証機能を実装" --label "Parent Task" --body "$(cat <<'EOF'
## 📋 概要

ユーザーがログイン・ログアウトできる認証機能を実装する。
セキュアな認証フローを提供し、ユーザー体験を向上させる。

## 📦 対象プロジェクト

- [x] 🔀 apis/bff
- [x] 🌐 clients/web

## ✅ 親タスク完了条件

- [ ] ログインページが作成され、未認証ユーザーのリダイレクト・ヘッダーからのログアウトが動作する
- [ ] Supabase Auth の Google OAuth でログインでき、認証後トップページへ遷移する
- [ ] BFF で認証トークンの検証が行われ、ログイン状態がページリロード後も維持される
- [ ] 認証エラー時に適切なエラーメッセージが表示される
- [ ] Supabase Auth が Terraform で IaC 化され、ローカルエミュレーターで開発できる
- [ ] 認証フローの E2E テストが CI で通過している
- [ ] staging 環境でのテストプレイが完了している

## 💬 備考

特に無し
EOF
)"
```

<!-- markdownlint-enable MD013 -->

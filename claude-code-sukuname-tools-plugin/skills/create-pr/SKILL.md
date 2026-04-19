---
name: create-pr
description: Pull Request を作成する。関連 Issue を紐づけてテンプレートに従って作成し、gh pr create で作成。
---

# create-pr スキル

Pull Request を作成する。

## 手順

1. 関連する Issue 番号をブランチ名から取得
2. 関連する Issue の内容とラベルを確認
3. 変更内容を `git diff main...HEAD` で確認
4. PR テンプレートに従って本文を作成
5. `gh pr create` で PR を作成（Issue と同じラベルを付与）

## 参照テンプレート

`.github/PULL_REQUEST_TEMPLATE.md`

## PR 作成実行の全体の例

```bash
gh pr create --title "#12 ユーザー認証APIを実装" --body "$(cat <<'EOF'
# Pull Request

## 📋 変更内容

### 📄 概要

ユーザー認証機能の API エンドポイントを実装した。

### 📂 変更ファイル

- `apis/auth/handler.go` - 認証ハンドラーを追加
- `apis/auth/service.go` - 認証サービスロジックを実装
- `apis/auth/repository.go` - ユーザーリポジトリを追加
- `apis/auth/middleware.go` - JWT 検証ミドルウェアを追加

### ➕ Issue 外の変更

無し

## ✅ 完了条件 (Issue に記載された内容)

- [x] POST /api/auth/login が Swagger 上で正常に動作確認が完了している
- [x] JWT ミドルウェアが実装され、手元での動作確認が完了している
- [x] 単体テスト・モックが用意されている
- [x] CI (test, lint, format) が全て正常に終了している

## 🏷️ タスク種別 (Issue に記載された内容)

- [ ] 🖥️ フロントエンド開発
- [x] ⚙️ バックエンド開発
- [ ] 🏗️ インフラ構築
- [ ] 📝 ドキュメント作成
- [ ] 🔧 設定
- [ ] ❓ その他

## 🖼️ スクリーンショット（UI 変更がある場合のみ）

無し

## 🔗 対応する Issue

Closes #12

## 💬 備考

認証トークンの有効期限は環境変数 `JWT_EXPIRY` で設定可能。
EOF
)" --label "Child Task" --base main
```

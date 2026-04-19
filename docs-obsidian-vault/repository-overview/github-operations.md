# GitHub 運用方針

Claude Code と GitHub を活用した開発手法。

## Issue

### 種類

- **Parent Task**: Epic/Feature 単位の大きなタスク (例: ユーザー認証機能を実装)
- **Child Task**: 親に紐づく個別タスク (例: ログイン API を実装、ログインページを作成)
- **todo**: シンプルな単発タスク (例: README を更新、軽微なミスを修正)
- **hotfix**: 緊急修正 (例: ログインエラーを修正)

### 親子関係

- Parent Task には完了条件を記載
- Child Task は親 Issue に紐づけ、親の完了条件を分解したタスク

```text
[Parent Task] ユーザー認証機能を実装
  ├── [Child Task] ログイン API を実装
  ├── [Child Task] ログインページを作成
  └── [Child Task] 認証ミドルウェアを追加
```

## ブランチ

- `feature/<issue番号>`: 通常の機能開発 (基本的にこちらを選択)
- `hotfix/<issue番号>`: 緊急修正

```bash
# 例: Issue #12 の機能開発
git checkout -b feature/12

# 例: Issue #42 の緊急修正
git checkout -b hotfix/42
```

## Pull Request

- ブランチ名から Issue 番号を取得して紐づけ
- `Closes #<issue番号>` でマージ時に自動クローズ
- Issue のラベル、完了条件を PR に引き継ぐ
- PR タイトル: `#<issue番号> <概要>` (例: `#12 ユーザー認証 API を実装`)

## 全体フロー

```text
[Issue 作成] -> [ブランチ作成] -> [実装] -> [コミット] -> [PR 作成] -> [レビュー] -> [マージ]
```

## Skills

各操作は Claude Code のスキルを使って実行すること。

### Issue 作成

- `/parent-issue`: Parent Task を作成
- `/child-issue`: Child Task を作成 (親 Issue に紐づけ)
- `/todo-issue`: シンプルな todo Issue を作成

### 開発フロー

- `/commit`: Conventional Commits 形式でコミット
- `/create-pr`: PR を作成 (Issue に紐づけ)
- `/apply-hotfix`: 全体フローを一括実行

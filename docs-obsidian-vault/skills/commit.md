---
name: commit
description: 現在の diff を確認してコミットメッセージを作成し、コミットを実行する。Conventional Commits 形式を使用。
---

# commit スキル

現在の diff を確認してコミットメッセージを作成し、コミットを実行する。

## 手順

1. `git status` で変更ファイルを確認
2. `git diff --staged` でステージ済みの変更を確認
3. ステージされていない場合は `git diff` で変更内容を確認
4. 変更内容に基づいてコミットメッセージを作成
5. Conventional Commits 形式でコミット

## コミットメッセージ形式

```text
<type>[optional scope]: #<issue番号> <description>

[optional body]
```

### type

- feat: 新機能
- fix: バグ修正
- docs: ドキュメント
- style: フォーマット
- refactor: リファクタリング
- test: テスト
- chore: ビルド・設定

### scope

- 変更対象の Nx プロジェクト名（例: `apis/bff`, `clients/web`）
- optional であるため、対象が多岐に渡る場合や Nx プロジェクト外の修正の場合は指定しない

### Issue 番号

- description の先頭に `#<issue番号>` 形式で記載
- ブランチ名から取得する（例: `feature/3` → `#3`）
- ブランチ名から特定できない場合は省略して良い

### body

- 指定条件無し
- optional であるため、小規模な変更でありタイトル以上に追加する情報がない場合は指定しない

## コミットメッセージ全体の例

### 最低限の例

```text
feat: #12 ログイン機能を追加
```

### 最大限の例

```text
refactor(clients/web): #34 コンポーネント構成を見直し

Atomic Design に基づいてコンポーネントを再構成。
Button, Input を atoms に、Form, Card を molecules に移動。
再利用性が向上し、依存関係が明確になった。
```

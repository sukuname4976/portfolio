---
name: git-commit
description: 現在のdiffを確認してコミットメッセージを作成し、コミットを実行する。Conventional Commits形式を使用。
---

# git-commit スキル

現在のdiffを確認してコミットメッセージを作成し、コミットを実行する。

## 手順

1. `git status` で変更ファイルを確認
2. `git diff --staged` でステージ済みの変更を確認
3. ステージされていない場合は `git diff` で変更内容を確認
4. 変更内容に基づいてコミットメッセージを作成
5. Conventional Commits形式でコミット

## コミットメッセージ形式

```
<type>(<scope>): <description>

<body>
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

変更対象のプロジェクト名やモジュール名

## 注意

- コミット前にユーザーに確認を取る
- 大きな変更は分割を提案する

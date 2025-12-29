---
name: create-pr
description: Pull Requestを作成する。関連Issueを紐づけてテンプレートに従って作成し、gh pr createで作成。
---

# create-pr スキル

Pull Requestを作成する。

## 手順

1. 関連するIssue番号を確認
2. 変更内容を `git diff main...HEAD` で確認
3. テンプレートに従ってPR本文を作成
4. `gh pr create` でPRを作成

## テンプレート

```markdown
## 関連Issue

Closes #xxx

## 概要

<!-- このPRで行った変更の概要 -->

## 変更内容

- 変更1
- 変更2

## テスト

- [ ] ローカルで動作確認済み
- [ ] テストを追加/更新済み

## スクリーンショット

<!-- UIの変更がある場合 -->

## レビュー観点

<!-- レビュアーに特に見てほしい点 -->
```

## コマンド例

```bash
gh pr create --title "タイトル" --body "本文" --base main
```

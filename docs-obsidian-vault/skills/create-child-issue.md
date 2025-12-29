---
name: create-child-issue
description: 子Issue（Task）を作成する。親Issueに紐づけてテンプレートに従って作成し、gh issue createで作成。
---

# create-child-issue スキル

子Issue（Task）を作成する。

## 手順

1. 親Issueの番号を確認
2. タスク内容をヒアリング
3. テンプレートに従ってIssue本文を作成
4. `gh issue create` でIssueを作成
5. 親Issueに子Issueへのリンクを追記

## テンプレート

```markdown
## 親Issue

#xxx

## 概要

<!-- このタスクで行うことの概要 -->

## 作業内容

- [ ] 作業1
- [ ] 作業2

## 完了条件

- [ ] 条件1
- [ ] 条件2

## 備考

<!-- 技術的な注意点、参考リンク等 -->
```

## コマンド例

```bash
gh issue create --title "タイトル" --body "本文" --label "task"
```

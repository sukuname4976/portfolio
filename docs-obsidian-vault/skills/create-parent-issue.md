---
name: create-parent-issue
description: 親Issue（Epic/Feature）を作成する。テンプレートに従ってIssue本文を作成し、gh issue createで作成。
---

# create-parent-issue スキル

親Issue（Epic/Feature）を作成する。

## 手順

1. ユーザーから要件をヒアリング
2. テンプレートに従ってIssue本文を作成
3. `gh issue create` でIssueを作成

## テンプレート

```markdown
## 概要

<!-- このIssueで達成したいことの概要 -->

## 背景・目的

<!-- なぜこの機能/改善が必要か -->

## 完了条件

- [ ] 条件1
- [ ] 条件2

## 子Issue

<!-- 作成後に追記 -->
- [ ] #xxx
- [ ] #xxx

## 備考

<!-- 参考リンク、関連Issue等 -->
```

## コマンド例

```bash
gh issue create --title "タイトル" --body "本文" --label "enhancement"
```

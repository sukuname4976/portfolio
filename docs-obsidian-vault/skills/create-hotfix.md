---
name: create-hotfix
description: hotfix を一括で実行する。Issue 作成 → 修正実装 → コミット → PR 作成まで自動化。
---

# create-hotfix スキル

hotfix を一括で実行する。Issue 作成から PR 作成まで自動化。

## 手順

1. ユーザーから hotfix 内容をヒアリング（何を修正するか）
2. `gh issue create` で hotfix Issue を作成
3. Issue 番号を取得し `hotfix/{issue番号}` ブランチを作成
4. AI が修正を実装
5. `git-commit` スキルを呼び出してコミット
6. `create-pr` スキルを呼び出して PR 作成（`Closes #{issue番号}`）

## Issue 作成実行の例

```bash
gh issue create --title "[Hotfix: ログインエラーを修正]" --label "hotfix" --body "$(cat <<'EOF'
## 問題

ログイン時に特定条件下でエラーが発生する。

## 修正方針

エラーハンドリングを追加し、適切なエラーメッセージを表示する。
EOF
)"
```

## ブランチ作成の例

```bash
# Issue 番号が #42 の場合
git checkout -b hotfix/42
```

## フロー全体

```text
[ヒアリング] → [Issue作成] → [ブランチ作成] → [修正実装] → [コミット] → [PR作成]
                   ↓              ↓              ↓            ↓           ↓
              hotfix ラベル  hotfix/{番号}    AI が実装   git-commit   create-pr
                                                           スキル      スキル
```

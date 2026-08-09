---
name: find-adr
description: 過去の意思決定 (ADR) を検索する。ある領域を変更する前や、方針を決める前に、既に決まっていることを確認して解き直しを防ぐ。
allowed-tools: Bash(*/skills/find-adr/scripts/list-adr.sh *)
---

# find-adr スキル

過去の意思決定 (ADR) を検索する。

ADR は `docs-obsidian-vault/**/decisions/` にあるが、エージェントの
コンテキストへ自動では読み込まれない。**触ろうとしている領域について
既に決まっていることがないかを、着手前に必ず確認する。**

## いつ使うか

- ある領域のコードや設定を変更する前
- 方針や技術選定を提案する前
- ADR を新しく書く前 (同じ課題を扱うものが既にないかの確認)

## 手順

1. 一覧を出す

   ```bash
   <skills>/find-adr/scripts/list-adr.sh
   ```

   1 行 1 件で `id` / `level` / タイトルが出る。本文は読まないため、
   件数が増えても出力は増えにくい。まず全件を眺めてよい。

2. 絞る

   ```bash
   # 置き場所で絞る
   <skills>/find-adr/scripts/list-adr.sh --dir infra-terraform

   # 階層で絞る (policy が最も上位)
   <skills>/find-adr/scripts/list-adr.sh --level policy
   ```

   タイトルで当たりが付かない場合は本文を検索する。

   ```bash
   grep -rl "<キーワード>" docs-obsidian-vault --include="*.md"
   ```

3. 読む

   上位から読む。`policy` がその範囲の前提、`design` が方法、
   `choice` は細かい取捨選択で、上位の方針には影響しない。

   急ぐ場合は `policy` と `design` だけ読めば骨格が掴める。

4. 無ければ、記録されていないと判断してよい

   一覧に該当が無ければ、その領域の決定は残っていない。探し続けず、
   新しく決めてよい。ただし選択肢を比べて片方を捨てたなら、
   `record-adr` で記録する。

## 読むときの注意

- 規範の正は ADR にある。README や `agents-base.md` と矛盾した場合は
  ADR に従う
- `却下事項` を必ず見る。そこに書かれている案を再び提案しないため
- `前提と見直し条件` を見る。前提が既に崩れているなら、その決定は
  見直しの対象になる

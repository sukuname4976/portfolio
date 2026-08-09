#!/usr/bin/env bash
# ADR の一覧を出力する。
# 使い方: list-adr.sh [--level <policy|design|choice>] [--dir <ディレクトリ名>]
#
# 1 行 1 件で id / level / タイトルを出力する。
# 本文は読まないため、件数が増えても出力量はタイトルの分しか増えない。
set -uo pipefail

level_filter=""
dir_filter=""

usage() {
  cat <<'USAGE'
使い方: list-adr.sh [--level <policy|design|choice>] [--dir <ディレクトリ名>]

  --level  階層で絞る (policy が最も上位)
  --dir    置き場所で絞る (例: infra-terraform、repository-overview)
USAGE
}

while [ "$#" -gt 0 ]; do
  case "$1" in
  --level)
    level_filter=${2:-}
    shift 2
    ;;
  --dir)
    dir_filter=${2:-}
    shift 2
    ;;
  -h | --help)
    usage
    exit 0
    ;;
  *)
    echo "不明な引数: $1" >&2
    usage >&2
    exit 2
    ;;
  esac
done

root=$(git rev-parse --show-toplevel 2>/dev/null) || root=$(pwd)
vault="${root}/docs-obsidian-vault"

if [ ! -d "$vault" ]; then
  echo "docs-obsidian-vault が見つからない: ${vault}" >&2
  exit 1
fi

count=0

while IFS= read -r file; do
  parent=$(basename "$(dirname "$(dirname "$file")")")
  [ -n "$dir_filter" ] && [ "$parent" != "$dir_filter" ] && continue

  frontmatter=$(sed -n '2,/^---$/p' "$file" | sed '$d')
  id=$(grep -m1 '^id:' <<<"$frontmatter" | sed 's/^id:[[:space:]]*//')
  level=$(grep -m1 '^level:' <<<"$frontmatter" | sed 's/^level:[[:space:]]*//')
  [ -n "$level_filter" ] && [ "$level" != "$level_filter" ] && continue

  title=$(grep -m1 '^# ' "$file" | sed "s/^# //; s/^${id}\. //")

  printf '%-28s %-8s %s\n' "$id" "$level" "$title"
  count=$((count + 1))
done < <(find "$vault" -path '*/decisions/*.md' -type f | sort)

if [ "$count" -eq 0 ]; then
  echo "該当する ADR は無い" >&2
else
  echo "-- ${count} 件"
fi

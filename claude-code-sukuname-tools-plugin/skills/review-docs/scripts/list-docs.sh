#!/usr/bin/env bash
# 追跡中の Markdown を 1 行 1 件で一覧する。
# 使い方: list-docs.sh [--grep <語>]
#
# パスと、h1 または「概要」の先頭行を出す。同じ主題を扱う文書を並べて、
# 記述が食い違っていないかを目で確かめるために使う。
set -uo pipefail

filter=""

usage() {
  cat <<'USAGE'
使い方: list-docs.sh [--grep <語>]

  --grep  パスまたは要約に語を含むものだけを出す
USAGE
}

while [ "$#" -gt 0 ]; do
  case "$1" in
  --grep)
    filter=${2:-}
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

root=$(git rev-parse --show-toplevel 2>/dev/null) || {
  echo "git リポジトリではない" >&2
  exit 1
}
cd "$root" || exit 1

count=0

while IFS= read -r file; do
  [ -f "$file" ] || continue

  # コードブロック内の見出しは例文なので除く
  heading=$(awk '
    /^```/ {fence = !fence; next}
    !fence && /^# / {sub(/^# /, ""); print; exit}
  ' "$file")
  overview=$(awk '
    /^```/ {fence = !fence; next}
    !fence && /^## 概要/ {found = 1; next}
    found && NF {print; exit}
  ' "$file")
  summary=${overview:-$heading}
  [ -n "$summary" ] || continue

  if [ -n "$filter" ]; then
    case "${file}|${summary}" in
    *"$filter"*) ;;
    *) continue ;;
    esac
  fi

  printf '%s\n    %s\n' "$file" "$summary"
  count=$((count + 1))
done < <(git ls-files '*.md' | sort)

if [ "$count" -eq 0 ]; then
  echo "該当する文書は無い" >&2
else
  echo "-- ${count} 件"
fi

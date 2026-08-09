#!/usr/bin/env bash
# ADR の構造を機械的に検証する。
# 使い方: check-adr.sh <ADR ファイル>...
#
# 内容の妥当性は判定しない。frontmatter・ID・ファイル名・節構成といった
# 形式だけを見る。用語の取り残しなど文脈に依存するものは検出できない。
set -uo pipefail

SECTIONS=(決定事項 コンテキスト 却下事項 前提と見直し条件 関連資料 変更履歴)
KEYS=(id summary status level created updated version)
LEVELS=(policy design choice)

status=0

report() {
  echo "$1: $2" >&2
  status=1
}

if [ "$#" -eq 0 ]; then
  echo "使い方: check-adr.sh <ADR ファイル>..." >&2
  exit 2
fi

for file in "$@"; do
  if [ ! -f "$file" ]; then
    report "$file" "ファイルが存在しない"
    continue
  fi

  if [ "$(head -n 1 "$file")" != "---" ]; then
    report "$file" "frontmatter が --- で始まっていない"
    continue
  fi

  frontmatter=$(sed -n '2,/^---$/p' "$file" | sed '$d')

  for key in "${KEYS[@]}"; do
    grep -q "^${key}:" <<<"$frontmatter" ||
      report "$file" "frontmatter に ${key} がない"
  done

  id=$(grep -m1 '^id:' <<<"$frontmatter" | sed 's/^id:[[:space:]]*//')
  version=$(grep -m1 '^version:' <<<"$frontmatter" | sed 's/^version:[[:space:]]*//')
  adr_status=$(grep -m1 '^status:' <<<"$frontmatter" | sed 's/^status:[[:space:]]*//')
  level=$(grep -m1 '^level:' <<<"$frontmatter" | sed 's/^level:[[:space:]]*//')

  case " ${LEVELS[*]} " in
  *" ${level} "*) ;;
  *) report "$file" "level が ${LEVELS[*]} のいずれでもない: ${level}" ;;
  esac

  grep -qE '^[0-9]+\.[0-9]+\.[0-9]+$' <<<"$version" ||
    report "$file" "version が <major>.<minor>.<patch> ではない: ${version}"

  if [ "$adr_status" != "採用" ] && [ "$adr_status" != "廃止" ]; then
    report "$file" "status が 採用 / 廃止 のいずれでもない: ${adr_status}"
  fi

  case "$(basename "$file" .md)" in
  "${id}-${level}-"*) ;;
  *) report "$file" "ファイル名が <id>-<level>- で始まっていない: ${id}-${level}" ;;
  esac

  prefix=${id%-*}
  parent=$(basename "$(dirname "$(dirname "$file")")")
  if [ "$prefix" != "$parent" ]; then
    report "$file" "id の接頭辞が置き場所と一致しない: ${prefix} / ${parent}"
  fi

  grep -q "^# ${id}\. " "$file" ||
    report "$file" "h1 が '# ${id}. ' の形式ではない"

  actual=$(grep '^## ' "$file" | sed 's/^## //')
  expected=$(printf '%s\n' "${SECTIONS[@]}")
  if [ "$actual" != "$expected" ]; then
    report "$file" "節の構成または順序が違う"
    diff <(echo "$expected") <(echo "$actual") | sed 's/^/    /' >&2
  fi
done

if [ "$status" -eq 0 ]; then
  echo "check-adr: 問題なし ($# 件)"
fi

exit "$status"

# Web Client

## 概要

Next.js 製の Web アプリ。フォーマットと lint は Biome、テストは Vitest、
BFF の型は openapi-typescript で生成する。

## 作業の進め方

- このファイルを読み込んだら、応答の冒頭に次の 1 行を出力する
  - `ドキュメントの指示に従い projects/client-web/agents-base.md を参照したことを通知します`
- 大きな変更をしたら `make prepare` を実行し、その時点で問題が無いことを
  確定させる。コミット時にまとめて手戻りが出るのを避けるため
- `src/api/generated/` は BFF の OpenAPI からの生成物なので直接編集しない。
  BFF 側のスキーマを変更して再生成する

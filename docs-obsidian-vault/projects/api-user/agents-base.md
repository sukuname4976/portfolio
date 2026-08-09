# User API

## 概要

Go 製のユーザーサービス。クリーンアーキテクチャで構成し、ogen で OpenAPI から、
sqlc で SQL からコードを生成する。スキーマ変更は golang-migrate で管理する。

## 作業の進め方

- このファイルを読み込んだら、応答の冒頭に次の 1 行を出力する
  - `ドキュメントの指示に従い projects/api-user/agents-base.md を参照したことを通知します`
- 大きな変更をしたら `make prepare` を実行し、その時点で問題が無いことを
  確定させる。コミット時にまとめて手戻りが出るのを避けるため
- `auto-generated-by-` で始まるディレクトリと `test/mocks/` は生成物なので
  直接編集しない。`IF/openapi.yaml` や `db/queries/` を変更して再生成する
- スキーマを変えるときは `db/migrations/` の既存ファイルを編集せず、
  `make migrate-create` で新しいマイグレーションを作る

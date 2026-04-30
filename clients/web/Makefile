.PHONY: dev build test lint lint-check format format-check gen-api gen-api-check check prepare clean

# サーバー起動 (開発用・ホットリロード)
dev:
	npm run dev

# ビルド (本番用)
build:
	npm run build

# サーバー起動 (本番ビルドのバイナリ実行)
up: build
	npm run start

# テスト
test:
	npm run test

# リンティング
lint:
	npm run lint

# リンティング確認
lint-check:
	npm run lint-check

# フォーマット
format:
	npm run format

# フォーマット確認
format-check:
	npm run format-check

# OpenAPI 仕様から TypeScript 型生成
gen-api:
	npm run gen-api

# 生成コードと OpenAPI spec の同期確認 (CI 検出用)
# tmpdir に生成 → 現行と比較し、差分があれば失敗とする
gen-api-check:
	@set -eu; \
	tmpfile=$$(mktemp); \
	trap 'rm -f "$$tmpfile"' EXIT; \
	npx openapi-typescript ../../apis/bff/IF/openapi.yaml -o "$$tmpfile" >/dev/null; \
	if ! diff -q "$$tmpfile" src/api/generated/schema.d.ts >/dev/null; then \
		echo "生成コードが OpenAPI spec と同期していません。"; \
		echo "make gen-api を実行してコミットしてください。"; \
		exit 1; \
	fi

# CI チェック (build + test + lint-check + format-check + gen-api-check)
check: build test lint-check format-check gen-api-check

# コミット前の品質担保処理 (build + test + lint + format + gen-api)
prepare: build test lint format gen-api

# 自動生成ファイルの削除
clean:
	rm -rf .next
	rm -rf src/api/generated

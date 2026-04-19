.PHONY: up dev gen-ogen gen-ogen-check gen-mock build test lint format format-check markdownlint check prepare clean

# サーバー起動 （本番同等のバイナリ実行）
up:
	./bin/user

# サーバー起動 （開発用・ホットリロード）
dev:
	PORT=8081 air -c .air.toml

# ogen で OpenAPI からコード生成
gen-ogen:
	ogen -config .ogen.yaml \
		-target src/presentation/auto-generated-by-ogen \
		-package ogen \
		-clean \
		IF/openapi.yaml

# mockgen でモック生成 (gateway-interfaces, repository-interfacesを自動検出)
gen-mock:
	@mkdir -p test/mocks
	@find src -name "*.go" \( -path "*/gateway-interfaces/*" -o -path "*/repository-interfaces/*" \) | while read file; do \
		dir=$$(dirname $$file); \
		name=$$(basename $$file .go); \
		pkg=$$(basename $$dir); \
		mockgen -source=$$file \
			-destination=test/mocks/$${pkg}_$${name}_mock.go \
			-package=mocks; \
		echo "Generated: test/mocks/$${pkg}_$${name}_mock.go"; \
	done

# ビルド
build:
	go build -o bin/user .

# テスト
test:
	go test ./... -v

# リンティング
lint:
	golangci-lint run --fix ./...

# リンティング確認
lint-check:
	golangci-lint run ./...

# フォーマット
format:
	goimports -w .
	gofmt -s -w .

# フォーマット確認
format-check:
	@test -z "$$(goimports -l .)" || (echo "以下のファイルのフォーマットまたはインポートが不正です:" && goimports -l . && exit 1)

# Markdown Lint 実行
markdownlint:
	npx markdownlint-cli2 "**/*.md"

# ogen 生成コードと OpenAPI spec の同期確認 (CI 検出用)
# tmpdir に生成 → 現行と比較し、差分があれば失敗とする
gen-ogen-check:
	@set -eu; \
	tmpdir=$$(mktemp -d); \
	trap 'rm -rf "$$tmpdir"' EXIT; \
	ogen -config .ogen.yaml -target "$$tmpdir" -package ogen -clean IF/openapi.yaml; \
	if ! diff -rq "$$tmpdir" src/presentation/auto-generated-by-ogen >/dev/null; then \
		echo "生成コードが OpenAPI spec と同期していません。"; \
		echo "make gen-ogen を実行してコミットしてください。"; \
		exit 1; \
	fi

# CI チェック (build + test + lint-check + format-check + gen-ogen-check)
# (Markdown Lint はリポジトリ全体を対象とし、node の依存関係が必要なため CI の別ジョブで実行する)
check: build test lint-check format-check gen-ogen-check

# コミット前の品質担保処理　　(build + test + lint + format + gen-ogen + gen-mock + markdownlint)
prepare: build test lint format gen-ogen gen-mock markdownlint

# 自動生成ファイルの削除
clean:
	rm -rf bin/
	rm -rf src/presentation/auto-generated-by-ogen/
	rm -rf test/mocks/

.PHONY: up dev gen-ogen gen-mock build test lint format format-check markdownlint check prepare clean

# サーバー起動 （本番同等のバイナリ実行）
up:
	./bin/bff

# サーバー起動 （開発用・ホットリロード）
dev:
	air -c .air.toml

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
	go build -o bin/bff .

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

#Markdown Lint 実行
markdownlint:
	npx markdownlint-cli2 "**/*.md"

# CI チェック (build + test + lint-check + format-check + markdownlint)
check: build test lint-check format-check markdownlint

# コミット前の品質担保処理　　(build + test + lint + format + gen-ogen + gen-mock + markdownlint)
prepare: build test lint format gen-ogen gen-mock markdownlint

# 自動生成ファイルの削除
clean:
	rm -rf bin/
	rm -rf src/presentation/auto-generated-by-ogen/
	rm -rf test/mocks/

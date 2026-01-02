.PHONY: generate generate-mock build test clean

# ogenでOpenAPIからコード生成
gen-openapi:
	ogen -config .ogen.yaml \
		-target src/presentation/auto-generated-by-ogen \
		-package ogen \
		-clean \
		IF/openapi.yaml

# mockgenでモック生成（gateway-interfaces, repository-interfacesを自動検出）
gen-mocks:
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

# テスト実行
test:
	go test ./... -v

# クリーン
clean:
	rm -rf bin/
	rm -rf src/presentation/auto-generated-by-ogen/

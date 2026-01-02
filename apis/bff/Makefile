.PHONY: generate build clean

# ogenでOpenAPIからコード生成
generate:
	ogen -config .ogen.yaml \
		-target src/presentation/auto-generated-by-ogen \
		-package ogen \
		-clean \
		IF/openapi.yaml

# ビルド
build:
	go build -o bin/bff .

# クリーン
clean:
	rm -rf bin/
	rm -rf src/presentation/auto-generated-by-ogen/

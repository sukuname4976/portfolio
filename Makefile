.PHONY: help install graph show-projects affected-graph

help:
	@echo "利用可能なコマンド:"
	@echo "  make install  - 依存関係をインストール"
	@echo "  make graph    - 依存グラフを表示"
	@echo "  make show     - プロジェクト一覧を表示"
	@echo "  make affected - 影響を受けるプロジェクトを表示"

# npm の依存関係をインストール
install:
	npm install

# Nx 関連コマンド
graph:
	npx nx graph

show:
	npx nx show projects

affected:
	npx nx affected:graph

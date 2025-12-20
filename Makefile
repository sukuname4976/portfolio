.PHONY: help install graph show-projects affected-graph

# デフォルトターゲット
help:
	@echo "Portfolio Monorepo - Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  help                      - Show this help message"
	@echo "  install                   - Install all dependencies"
	@echo "  graph                     - Show dependency graph in browser"
	@echo "  show-projects             - List all projects"
	@echo "  affected-graph            - Show affected projects graph"
	@echo ""
	@echo "Note: Project-specific operations should be done in each project directory"
	@echo "      or using Nx commands (e.g., 'nx build bff', 'nx serve web')"

# 依存関係のインストール
install:
	@echo "Installing dependencies..."
	npm install

# 依存関係グラフを表示
graph:
	@echo "Opening dependency graph in browser..."
	npx nx graph

# プロジェクト一覧を表示
show-projects:
	@echo "Listing all projects..."
	npx nx show projects

# 影響を受けるプロジェクトのグラフを表示
affected-graph:
	@echo "Opening affected projects graph in browser..."
	npx nx affected:graph


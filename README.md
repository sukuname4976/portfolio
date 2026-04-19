# Portfolio Monorepo

このリポジトリは、ポートフォリオプロジェクトの monorepoです。
このプロジェクトは、プレーリーカード等で初めて会う人に共有する自己紹介や制作物・スキルの紹介を行うアプリケーションを提供します。

大きく以下の3つの実装を統合管理します。

- バックエンドの API マイクロサービス
- フロントエンドのクライアントアプリケーション
- Infrastructure as Code（Terraform）

ただし、現状として infras/, libraries/ はディレクトリが存在しているのみです。
単なる拡張予定として基本的に無視してください。

## 📁 ディレクトリ構造

```text
portfolio/
├── .claude/                          # Claude Code 設定
│   ├── settings.json                   # Claude Code 設定ファイル
│   └── skills/                         # monorepoルート向けカスタムスキル定義
├── .devcontainer/                    # 言語別DevContainer設定
│   ├── go/                             # Go 開発環境
│   ├── python/                         # Python 開発環境
│   ├── nodejs/                         # Node.js/Next.js 開発環境
│   └── terraform/                      # Terraform 開発環境
├── .github/                          # GitHub設定
│   ├── ISSUE_TEMPLATE/                 # Issueテンプレート
│   ├── PULL_REQUEST_TEMPLATE.md        # Pull Requestテンプレート
│   └── workflows/                      # GitHub Actions CI/CD
├── apis/                             # バックエンドAPIマイクロサービス
│   ├── bff/                            # Go - Backend for Frontendサービス（ポート: 8080）
│   ├── user/                           # Go - Userサービス（ポート: 8081）
│   └── agent/                          # Python - Agentサービス
├── clients/                          # フロントエンドクライアント
│   └── web/                            # Next.js
├── docs-obsidian-vault/              # ドキュメント（Obsidian Vault）
│   ├── projects/                       # プロジェクト別ドキュメント
│   ├── repository-overview/            # リポジトリ全体のドキュメント
│   └── skills/                         # スキル定義ドキュメント
├── infras/                           # Infrastructure as Code
│   └── terraform/                      # Terraform
└── libraries/                        # 共有ライブラリ
    ├── golang-shared/                  # Go 共通ライブラリ
    └── python-shared/                  # Python 共通ライブラリ
```

## リポジトリルートでの共通化

本来monorepoの各プロジェクトは外部に依存せず独立して動作するのが理想です。
しかし、個人開発の効率化のためにこのリポジトリではルートで以下を共通化しています。

- Nx（monorepo 管理、依存関係グラフ、差分検出）
- markdownlint-cli2（Markdown のリンティング）
- husky（pre-commit フック）
- GitHub Actions CI/CD
- DevContainer（各言語の開発環境）
- docs-obsidian-vault（プロンプト・ドキュメントの統括管理）

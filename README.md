# Portfolio Monorepo

このリポジトリは、ポートフォリオプロジェクトのmonorepoです。バックエンドのマイクロサービス、フロントエンドアプリケーション（Web、Android、iOS）、およびInfrastructure as Code（Terraform）を統合管理します。

## 📁 ディレクトリ構造

```
portfolio/
├── .devcontainer/          # 言語別DevContainer設定
│   ├── go/                 # Go開発環境
│   ├── python/             # Python開発環境
│   ├── nodejs/             # Node.js/Next.js開発環境
│   ├── kotlin/             # Kotlin/Android開発環境
│   ├── swift/              # Swift/iOS開発環境
│   └── terraform/          # Terraform開発環境
├── apis/                   # バックエンドAPIサービス
│   ├── bff/               # Go - Backend for Frontend（ポート: 8080）
│   ├── user/              # Go - ユーザーサービス（ポート: 8081）
│   └── agent/             # Python - Agentサービス
├── clients/                # フロントエンドクライアントアプリケーション
│   ├── web/               # Next.js
│   ├── android/           # Kotlin
│   └── ios/               # Swift
└── infras/                 # Infrastructure as Code
    └── terraform/         # Terraform
```

## 🚀 セットアップ

### 前提条件

- Docker Desktop（またはDocker + Docker Compose）
- Visual Studio Code
- Dev Containers拡張機能

### 開発環境の起動

1. リポジトリをクローン
```bash
git clone https://github.com/sukuname4976/portfolio.git
cd portfolio
```

2. VS Codeでプロジェクトを開く
```bash
code .
```

3. Dev Containerで開く
   - 各プロジェクト（例: `apis/bff`）を開く
   - `F1` → `Dev Containers: Reopen in Container`を選択
   - または、VS Codeの通知から「コンテナで再度開く」をクリック

## 🛠️ 各サービスの開発

### Goサービス（bff, user）

```bash
# bffサービスの開発
cd apis/bff
# DevContainerで開く（.devcontainer.jsonが自動的に../../.devcontainer/go/devcontainer.jsonを参照）

# サービス起動
nx serve bff
# または
go run main.go
```

### Pythonサービス（agent）

```bash
# agentサービスの開発
cd apis/agent
# DevContainerで開く

# サービス起動
nx serve agent
# または
python main.py
```

### Next.jsアプリ（web）

```bash
# webアプリの開発
cd clients/web
# DevContainerで開く

# 開発サーバー起動
nx serve web
# または
npm run dev
```

### Androidアプリ

```bash
# androidアプリの開発
cd clients/android
# DevContainerで開く

# ビルド
nx build android
```

### iOSアプリ

```bash
# iosアプリの開発
cd clients/ios
# 注意: 完全なiOS開発にはmacOSとXcodeが必要です
```

### Terraform

```bash
# Terraformの開発
cd infras/terraform
# DevContainerで開く

# 初期化
nx init terraform
# プラン
nx plan terraform
```

## 📚 Nxコマンドの使用

このプロジェクトはNxで管理されています。

```bash
# 影響を受けたプロジェクトをビルド
nx affected -t build

# 影響を受けたプロジェクトをテスト
nx affected -t test

# 依存関係グラフを表示
nx graph

# 特定のプロジェクトをビルド
nx build <project-name>

# 特定のプロジェクトを起動
nx serve <project-name>
```

## 🔧 DevContainerの構成

各言語用のDevContainerは`.devcontainer/`配下に配置されています：

- `.devcontainer/go/` - Go開発環境（bff, userで共有）
- `.devcontainer/python/` - Python開発環境（agent用）
- `.devcontainer/nodejs/` - Node.js/Next.js開発環境（web用）
- `.devcontainer/kotlin/` - Kotlin/Android開発環境
- `.devcontainer/swift/` - Swift/iOS開発環境
- `.devcontainer/terraform/` - Terraform開発環境

各プロジェクトは`.devcontainer.json`でルートのDevContainerを参照します。

## 📝 注意事項

- Goサービス（`bff`, `user`）は同じGo DevContainerを共有します
- ポートは環境変数で管理し、競合を避けてください
- iOS開発はWSL2では制限があるため、macOSでの開発を推奨します
- 各サービスは独立して動作確認可能です

## 🤝 コントリビューション

プルリクエストを歓迎します。大きな変更の場合は、まずイシューを開いて変更内容を議論してください。

## 🤖 Claude Code（GitHub Actions）による自動レビュー

Pull Request作成/更新時に、Claudeが自動でコードレビューコメントを投稿します。

### セットアップ（リポジトリ管理者）

- **Secretsの追加**: GitHub リポジトリの `Settings > Secrets and variables > Actions` で以下を登録します。
  - **`CLAUDE_CODE_OAUTH_TOKEN`**: Claude Code（Proアカウント）用のOAuthトークン

ワークフローは `.github/workflows/claude-code-PR-review.yml` に定義されています。

### プロンプトの調整

- プロンプトは `.github/workflows/claude-review.md` を編集して調整します（PRごとに自動で反映されます）。

### 注意

- **認証情報の取り扱い**: `CLAUDE_CODE_OAUTH_TOKEN` は機密情報です。GitHub Secretsで管理し、必要に応じて更新してください。

## 📄 ライセンス

[ライセンス情報を追加]


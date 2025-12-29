# Portfolio Monorepo

このリポジトリは、ポートフォリオプロジェクトのmonorepoです。バックエンドのマイクロサービス、フロントエンドアプリケーション（Web、Android、iOS）、およびInfrastructure as Code（Terraform）を統合管理します。

現在大幅なリポジトリ構造の編集中であり、完了後に手動で更新されるまではREADME.mdの編集はしません。信頼も置いてはいけません。

## 📁 ディレクトリ構造

```text
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

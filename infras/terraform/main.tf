terraform {
  required_version = ">= 1.0"

  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6.0"
    }
  }

  # state はローカル管理（手動ローカル apply 前提）。
  # リモート化する場合はここに backend を定義する。
  # backend "s3" {
  #   bucket = "your-terraform-state-bucket"
  #   key    = "portfolio/terraform.tfstate"
  #   region = "us-east-1"
  # }
}

provider "github" {
  owner = var.github_owner

  # 認証は環境変数 GITHUB_TOKEN から自動取得する。
  # `source ./load-env.sh` で `export GITHUB_TOKEN=$(gh auth token)` を注入してから実行する。
  # トークンの値はコードにもファイルにも保存しない。
}

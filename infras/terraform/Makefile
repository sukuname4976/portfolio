.PHONY: init plan apply fmt fmt-check init-check validate check prepare

# 初期化 (provider の取得)
init:
	terraform init

# 差分確認
plan:
	terraform plan

# 適用
# 運用ルール上、変更をコミットしてから実行する (README.md を参照)
apply:
	terraform apply

# フォーマット
fmt:
	terraform fmt -recursive

# フォーマット確認
fmt-check:
	terraform fmt -check -recursive

# 検証用の初期化
# -backend=false     : backend 設定なしで初期化する
# -lockfile=readonly : ロックファイルを書き換えず、古ければ失敗させる
# GitHub API を叩かないため認証情報は不要
init-check:
	terraform init -backend=false -lockfile=readonly

# 構文·型·参照の検証
validate: init-check
	terraform validate

# CI チェック (フォーマット確認 + 検証)
check: fmt-check validate

# コミット前の品質担保処理 (フォーマット + 検証)
prepare: fmt validate

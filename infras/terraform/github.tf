# GitHub リポジトリの設定を IaC 化する。
# 対象は既存リポジトリの「保護設定」のみで、リポジトリ自体は Terraform 管理下に置かない。

# main ブランチの保護。
# 従来の branch protection ではなく ruleset を使う（GitHub が推奨する後継の仕組み）。
resource "github_repository_ruleset" "main_protection" {
  name        = "main-protection"
  repository  = var.repository_name
  target      = "branch"
  enforcement = "active"

  conditions {
    ref_name {
      # ~DEFAULT_BRANCH はデフォルトブランチ（= main）を指す組み込みパターン。
      # ブランチ名を直書きしないことで、デフォルトブランチを変えても追従する。
      include = ["~DEFAULT_BRANCH"]
      exclude = []
    }
  }

  rules {
    # ブランチ削除の禁止。
    deletion = true

    # force-push の禁止。過去に main を force-push で壊した事故の再発防止。
    non_fast_forward = true

    # 直 push を禁止し、変更は必ず PR 経由にする。
    pull_request {
      # ソロ運用のため 0。GitHub は自分の PR を自分で承認できないので、
      # 1 以上にすると自分の PR がマージ不能になり詰む。
      # ここでの目的は「PR を経由させること（直 push の禁止）」なので 0 で足りる。
      required_approving_review_count = 0

      dismiss_stale_reviews_on_push     = false
      require_code_owner_review         = false
      require_last_push_approval        = false
      required_review_thread_resolution = false
    }

    # CI の成功をマージの条件にする。
    # PR を経由させるだけでは CI が落ちていてもマージできてしまうため、
    # .github/workflows/ci.yaml の 2 ジョブが成功しない限りマージできないようにする。
    required_status_checks {
      # context はワークフローのジョブ名 (jobs.<id>.name) と一致させる。
      # integration_id は指定しない。GitHub Actions のチェックは省略しても対象になる。
      required_check {
        context = "Nx 変更検出したプロジェクトへ CI を実行"
      }

      required_check {
        context = "リポジトリ全体へ Markdown Lint を実行"
      }

      # strict を true にすると、PR のブランチが main の最新を取り込んでいないと
      # マージできなくなる。ソロ運用では並行して進む PR が少なく、main が進むたびに
      # PR ごとに取り込み直す手間が見合わないため false のままにする。
      strict_required_status_checks_policy = false
    }
  }

  # bypass_actors は意図的に定義しない。
  # admin（= 本人）も直 push できない状態を「最後の砦」として維持する。
  # 緊急時のバイパスが必要になった場合にのみ追加を検討する。
}

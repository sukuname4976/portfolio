variable "github_owner" {
  description = "GitHub のオーナー（ユーザーまたは Organization）"
  type        = string
  default     = "sukuname4976"
}

variable "repository_name" {
  description = "設定対象の GitHub リポジトリ名"
  type        = string
  default     = "portfolio"
}

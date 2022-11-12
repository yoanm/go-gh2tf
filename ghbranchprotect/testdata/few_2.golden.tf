resource "github_branch_protection" "branch-protection-id" {
  repository_id           = "repository_name"
  pattern                 = "branch-pattern"
  enforce_admins          = true
  allows_force_pushes     = true
  required_linear_history = false
  require_signed_commits  = true

  required_status_checks {
    strict = true
  }

  required_pull_request_reviews {
    dismiss_stale_reviews           = true
    require_code_owner_reviews      = true
    required_approving_review_count = 2
  }
}

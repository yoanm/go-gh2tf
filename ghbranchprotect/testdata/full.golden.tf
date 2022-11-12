resource "github_branch_protection" "branch-protection-id" {
  repository_id           = "repository_name"
  pattern                 = "branch-pattern"
  enforce_admins          = true
  allows_deletions        = false
  allows_force_pushes     = true
  push_restrictions       = ["push-restrictions"]
  required_linear_history = false
  require_signed_commits  = true

  required_status_checks {
    strict   = false
    contexts = ["context"]
  }

  required_pull_request_reviews {
    dismiss_stale_reviews           = true
    restrict_dismissals             = false
    dismissal_restrictions          = ["dismissal-restriction"]
    require_code_owner_reviews      = true
    required_approving_review_count = 2
  }
}

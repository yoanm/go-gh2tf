resource "github_branch_protection" "branch-protection-id" {
  repository_id          = "repository_name"
  pattern                = "branch-pattern"
  allows_deletions       = false
  push_restrictions      = ["push-restrictions"]
  require_signed_commits = true

  required_status_checks {
    contexts = ["context"]
  }

  required_pull_request_reviews {
    restrict_dismissals    = true
    dismissal_restrictions = ["dismissal-restriction"]
  }
}

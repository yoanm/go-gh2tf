resource "github_branch_protection" "branch-protection-id" {
  repository_id           = "repository_name"
  pattern                 = "branch-pattern"
  enforce_admins          = true
  allows_deletions        = false
  allows_force_pushes     = true
  push_restrictions       = ["push-restrictions"]
  required_linear_history = false
  require_signed_commits  = true
}

resource "github_branch" "branch-id" {
  repository    = "repository_name"
  branch        = "my_branch_name"
  source_branch = "source_branch"
  source_sha    = "source_sha"
}

resource "github_repository" "repository-id" {
  name = "repository_name"

  visibility  = "visibility"
  description = "description"

  topics = ["topic"]

  has_issues    = true
  has_downloads = true

  delete_branch_on_merge = true

  vulnerability_alerts = true

  archive_on_destroy = true
}

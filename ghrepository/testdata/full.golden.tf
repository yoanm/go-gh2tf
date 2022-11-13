resource "github_repository" "repository-id" {
  name      = "repository_name"
  auto_init = false

  visibility  = "visibility"
  description = "description"

  template {
    owner      = "owner"
    repository = "repository"
  }

  topics       = ["topic"]
  homepage_url = "homepage-url"

  pages {
    source {
      branch = "branch"
      path   = "source"
    }
  }

  has_issues    = true
  has_projects  = false
  has_wiki      = true
  has_downloads = false

  allow_merge_commit     = false
  allow_rebase_merge     = true
  allow_squash_merge     = false
  allow_auto_merge       = true
  delete_branch_on_merge = false

  merge_commit_title          = "merge-commit-title"
  merge_commit_message        = "merge-commit-message"
  squash_merge_commit_title   = "squash-merge-commit-title"
  squash_merge_commit_message = "squash-merge-commit-message"

  vulnerability_alerts = true

  archived           = false
  archive_on_destroy = true
}

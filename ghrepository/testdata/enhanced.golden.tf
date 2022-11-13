resource "github_repository" "repository-id" {
  name = "repository_name"

  visibility = "visibility"

  template {
    owner      = "owner"
    repository = "repository"
  }

  homepage_url = "homepage-url"

  pages {
    source {
      branch = "branch"
    }
  }

  has_projects = true
  has_wiki     = true

  allow_rebase_merge = true
  allow_auto_merge   = true

  merge_commit_title        = "merge-commit-title"
  squash_merge_commit_title = "squash-merge-commit-title"

  archive_on_destroy = true
}

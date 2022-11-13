package ghrepository_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-gh2tf"
	"github.com/yoanm/go-gh2tf/ghrepository"
)

func ExampleNew() {
	valGen := gh2tf.NewValueGenerator()
	repoName := "repository_name"
	visibility := "var.default_visibility"
	trueBool := "true"
	falseBool := "false"
	topics := []string{"topic"}
	description := "local.my_local"
	homepageUrl := "homepage-url"
	mergeCommitTitle := "merge-commit-title"
	mergeCommitMessage := "merge-commit-message"
	squashMergeCommitTitle := "squash-merge-commit-title"
	squashMergeCommitMessage := "squash-merge-commit-message"
	repoPageSourcePath := "branch"
	repoPageSourceSource := "source"
	templateOwner := "owner"
	templateRepository := "repository"
	res := &ghrepository.Config{
		valGen,
		"repository-id",
		&repoName,
		&visibility,
		&falseBool,
		&description,
		&falseBool,
		&trueBool,
		&falseBool,
		&trueBool,
		&falseBool,
		&homepageUrl,
		&topics,

		&trueBool,
		&falseBool,
		&trueBool,
		&falseBool,
		&trueBool,
		&mergeCommitTitle,
		&mergeCommitMessage,
		&squashMergeCommitTitle,
		&squashMergeCommitMessage,
		&falseBool,
		&trueBool,
		&ghrepository.PagesConfig{
			&ghrepository.PagesSourceConfig{valGen, &repoPageSourcePath, &repoPageSourceSource},
		},
		&ghrepository.TemplateConfig{valGen, &templateOwner, &templateRepository},
	}

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendBlock(ghrepository.New(res))
	fmt.Println(string(hclFile.Bytes()))

	// Output:
	// resource "github_repository" "repository-id" {
	//   name      = "repository_name"
	//   auto_init = false
	//
	//   visibility  = var.default_visibility
	//   description = local.my_local
	//
	//   template {
	//     owner      = "owner"
	//     repository = "repository"
	//   }
	//
	//   topics       = ["topic"]
	//   homepage_url = "homepage-url"
	//
	//   pages {
	//     source {
	//       branch = "branch"
	//       path   = "source"
	//     }
	//   }
	//
	//   has_issues    = true
	//   has_projects  = false
	//   has_wiki      = true
	//   has_downloads = false
	//
	//   allow_merge_commit     = false
	//   allow_rebase_merge     = true
	//   allow_squash_merge     = false
	//   allow_auto_merge       = true
	//   delete_branch_on_merge = false
	//
	//   merge_commit_title          = "merge-commit-title"
	//   merge_commit_message        = "merge-commit-message"
	//   squash_merge_commit_title   = "squash-merge-commit-title"
	//   squash_merge_commit_message = "squash-merge-commit-message"
	//
	//   vulnerability_alerts = true
	//
	//   archived           = false
	//   archive_on_destroy = true
	// }
}

package ghbranch_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/yoanm/go-gh2tf"
	"github.com/yoanm/go-gh2tf/ghbranch"
)

func ExampleNew() {
	valGen := gh2tf.NewValueGenerator()
	repoName := "repository_name"
	branchName := "my_default_branch_name"
	sourceBranch := "source_branch_name"
	sourceSha := "a_sha"
	repoNameAsLink := "github_repository.res-id.name"
	branchNameAsLink := "github_branch.res-id.branch"

	res := &ghbranch.Config{
		valGen,
		"res-id",
		&repoName,
		&branchName,
		nil,
		nil,
	}
	resWithSourceBranch := &ghbranch.Config{
		valGen,
		"res-with-source-branch",
		&repoNameAsLink,
		&branchName,
		&sourceBranch,
		nil,
	}
	resWithSourceSha := &ghbranch.Config{
		valGen,
		"res-with-source-sha",
		&repoName,
		&branchNameAsLink,
		nil,
		&sourceSha,
	}

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendBlock(ghbranch.New(res))
	hclFile.Body().AppendBlock(ghbranch.New(resWithSourceBranch))
	hclFile.Body().AppendBlock(ghbranch.New(resWithSourceSha))
	fmt.Println(string(hclFile.Bytes()))

	// Output:
	// resource "github_branch" "res-id" {
	//   repository = "repository_name"
	//   branch     = "my_default_branch_name"
	// }
	// resource "github_branch" "res-with-source-branch" {
	//   repository    = github_repository.res-id.name
	//   branch        = "my_default_branch_name"
	//   source_branch = "source_branch_name"
	// }
	// resource "github_branch" "res-with-source-sha" {
	//   repository = "repository_name"
	//   branch     = github_branch.res-id.branch
	//   source_sha = "a_sha"
	// }
}

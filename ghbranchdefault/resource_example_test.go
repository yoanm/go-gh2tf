package ghbranchdefault_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/yoanm/go-gh2tf"
	"github.com/yoanm/go-gh2tf/ghbranchdefault"
)

func ExampleNew() {
	valGen := gh2tf.NewValueGenerator()
	repoName := "repository_name"
	branchName := "my_default_branch_name"

	res := &ghbranchdefault.Config{
		valGen,
		"res-id",
		&repoName,
		&branchName,
	}
	repoNameAsLink := "github_repository.res-id.name"
	branchNameAsLink := "github_branch.res-id.branch"

	resWithLinks := &ghbranchdefault.Config{
		valGen,
		"res-id-with-links",
		&repoNameAsLink,
		&branchNameAsLink,
	}

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendBlock(ghbranchdefault.New(res))
	hclFile.Body().AppendBlock(ghbranchdefault.New(resWithLinks))
	fmt.Println(string(hclFile.Bytes()))

	// Output:
	// resource "github_branch_default" "res-id" {
	//   repository = "repository_name"
	//   branch     = "my_default_branch_name"
	// }
	// resource "github_branch_default" "res-id-with-links" {
	//   repository = github_repository.res-id.name
	//   branch     = github_branch.res-id.branch
	// }
}

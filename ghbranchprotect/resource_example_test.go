package ghbranchprotect

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/yoanm/go-gh2tf"
)

func ExampleNew() {
	valGen := gh2tf.NewValueGenerator()
	repoName := "repository_name"
	branchPattern := "branch-pattern"
	trueBool := "true"
	falseBool := "false"
	pushRestrictions := []string{"data.github_teams.my_team.node_id"}
	contexts := []string{"context"}
	dismissalRestrictions := []string{"dismissal-restriction"}
	two := "2"
	res := &Config{
		valGen,
		"branch-protection-id",
		&repoName,
		&branchPattern,
		&trueBool,
		&falseBool,
		&trueBool,
		&pushRestrictions,
		&falseBool,
		&trueBool,

		&RequiredStatusChecksConfig{
			valGen,
			&falseBool,
			&contexts,
		},

		&RequiredPRReviewConfig{
			valGen,
			&trueBool,
			&falseBool,
			&dismissalRestrictions,
			&trueBool,
			&two,
		},
	}
	hclBlock := New(res)

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendBlock(hclBlock)
	fmt.Println(string(hclFile.Bytes()))

	// Output:
	// resource "github_branch_protection" "branch-protection-id" {
	//   repository_id           = "repository_name"
	//   pattern                 = "branch-pattern"
	//   enforce_admins          = true
	//   allows_deletions        = false
	//   allows_force_pushes     = true
	//   push_restrictions       = [data.github_teams.my_team.node_id]
	//   required_linear_history = false
	//   require_signed_commits  = true
	//
	//   required_status_checks {
	//     strict   = false
	//     contexts = ["context"]
	//   }
	//
	//   required_pull_request_reviews {
	//     dismiss_stale_reviews           = true
	//     restrict_dismissals             = false
	//     dismissal_restrictions          = ["dismissal-restriction"]
	//     require_code_owner_reviews      = true
	//     required_approving_review_count = 2
	//   }
	// }
}

package ghbranch_test

import (
	"testing"

	"github.com/yoanm/go-gh2tf"
	"github.com/yoanm/go-gh2tf/ghbranch"
	"github.com/yoanm/go-tfsig/testutils"
)

func TestNew(t *testing.T) {
	t.Parallel()

	valGen := gh2tf.NewValueGenerator()
	repoName := "repository_name"
	branchName := "my_branch_name"
	sourceBranch := "source_branch"
	sourceSha := "source_sha"
	cases := map[string]struct {
		value      *ghbranch.Config
		goldenFile string
	}{
		"Basic": {
			&ghbranch.Config{
				valGen,
				"branch-id",
				&repoName,
				&branchName,
				nil,
				nil,
			},
			"base",
		},
		"Full": {
			&ghbranch.Config{
				valGen,
				"branch-id",
				&repoName,
				&branchName,
				&sourceBranch,
				&sourceSha,
			},
			"full",
		},
		"Has not resource": {
			&ghbranch.Config{valGen, "", nil, nil, nil, nil},
			"empty",
		},
		"Nil config": {
			nil,
			"empty",
		},
	}

	for tcname, tcase := range cases {
		tcase := tcase // For parallel execution

		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				if err := testutils.EnsureBlockFileEqualsGoldenFile(ghbranch.New(tcase.value), tcase.goldenFile); err != nil {
					t.Errorf("Case \"%s\": %v", t.Name(), err)
				}
			},
		)
	}
}

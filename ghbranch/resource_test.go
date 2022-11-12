package ghbranch

import (
	"testing"

	"github.com/yoanm/go-tfsig/testutils"

	"github.com/yoanm/go-gh2tf"
)

func TestNew(t *testing.T) {
	valGen := gh2tf.NewValueGenerator()
	repoName := "repository_name"
	branchName := "my_branch_name"
	sourceBranch := "source_branch"
	sourceSha := "source_sha"
	cases := map[string]struct {
		value      *Config
		goldenFile string
	}{
		"Basic": {
			&Config{
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
			&Config{
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
			&Config{valGen, "", nil, nil, nil, nil},
			"empty",
		},
		"Nil config": {
			nil,
			"empty",
		},
	}

	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				if err := testutils.EnsureBlockFileEqualsGoldenFile(New(tc.value), tc.goldenFile); err != nil {
					t.Errorf("Case \"%s\": %v", tcname, err)
				}
			},
		)
	}
}

package ghbranchdefault_test

import (
	"testing"

	"github.com/yoanm/go-tfsig/testutils"

	"github.com/yoanm/go-gh2tf"
	"github.com/yoanm/go-gh2tf/ghbranchdefault"
)

func TestNew(t *testing.T) {
	t.Parallel()

	valGen := gh2tf.NewValueGenerator()
	repoName := "repository_name"
	branchName := "my_branch_name"
	cases := map[string]struct {
		value      *ghbranchdefault.Config
		goldenFile string
	}{
		"Basic": {
			&ghbranchdefault.Config{
				valGen,
				"default-branch-id",
				&repoName,
				&branchName,
			},
			"base",
		},
		"Has not resource": {
			&ghbranchdefault.Config{valGen, "", nil, nil},
			"empty",
		},
		"Nil config": {
			nil,
			"empty",
		},
	}

	for tcname, tcase := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				actual := ghbranchdefault.New(tcase.value)
				if err := testutils.EnsureBlockFileEqualsGoldenFile(actual, tcase.goldenFile); err != nil {
					t.Errorf("Case \"%s\": %v", t.Name(), err)
				}
			},
		)
	}
}

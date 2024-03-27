package ghbranchprotect_test

import (
	"testing"

	"github.com/yoanm/go-tfsig/testutils"

	"github.com/yoanm/go-gh2tf"
	"github.com/yoanm/go-gh2tf/ghbranchprotect"
)

func TestNew(t *testing.T) {
	t.Parallel()

	valGen := gh2tf.NewValueGenerator()
	repoName := "repository_name"
	branchPattern := "branch-pattern"
	trueBool := "true"
	falseBool := "false"
	pushRestrictions := []string{"push-restrictions"}
	contexts := []string{"context"}
	dismissalRestrictions := []string{"dismissal-restriction"}
	two := "2"
	cases := map[string]struct {
		value      *ghbranchprotect.Config
		goldenFile string
	}{
		"Full": {
			&ghbranchprotect.Config{
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

				&ghbranchprotect.RequiredStatusChecksConfig{
					valGen,
					&falseBool,
					&contexts,
				},

				&ghbranchprotect.RequiredPRReviewConfig{
					valGen,
					&trueBool,
					&falseBool,
					&dismissalRestrictions,
					&trueBool,
					&two,
				},
			},
			"full",
		},
		"Short": {
			&ghbranchprotect.Config{
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

				nil,
				nil,
			},
			"short",
		},
		"Few #1": {
			&ghbranchprotect.Config{
				valGen,
				"branch-protection-id",
				&repoName,
				&branchPattern,
				nil,
				&falseBool,
				nil,
				&pushRestrictions,
				nil,
				&trueBool,

				&ghbranchprotect.RequiredStatusChecksConfig{
					valGen,
					nil,
					&contexts,
				},

				&ghbranchprotect.RequiredPRReviewConfig{
					valGen,
					nil,
					&trueBool,
					&dismissalRestrictions,
					nil,
					nil,
				},
			},
			"few_1",
		},

		"Few #2": {
			&ghbranchprotect.Config{
				valGen,
				"branch-protection-id",
				&repoName,
				&branchPattern,
				&trueBool,
				nil,
				&trueBool,
				nil,
				&falseBool,
				&trueBool,

				&ghbranchprotect.RequiredStatusChecksConfig{
					valGen,
					&trueBool,
					nil,
				},

				&ghbranchprotect.RequiredPRReviewConfig{
					valGen,
					&trueBool,
					nil,
					nil,
					&trueBool,
					&two,
				},
			},
			"few_2",
		},
		"Has not resource": {
			&ghbranchprotect.Config{valGen, "branch-protection-id", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
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

				actual := ghbranchprotect.New(tcase.value)
				if err := testutils.EnsureBlockFileEqualsGoldenFile(actual, tcase.goldenFile); err != nil {
					t.Errorf("Case \"%s\": %v", t.Name(), err)
				}
			},
		)
	}
}

package ghbranchprotect

import (
	"testing"

	"github.com/yoanm/go-tfsig/testutils"

	"github.com/yoanm/go-gh2tf"
)

func TestNew(t *testing.T) {
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
		value      *Config
		goldenFile string
	}{
		"Full": {
			&Config{
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
			},
			"full",
		},
		"Short": {
			&Config{
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
			&Config{
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

				&RequiredStatusChecksConfig{
					valGen,
					nil,
					&contexts,
				},

				&RequiredPRReviewConfig{
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
			&Config{
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

				&RequiredStatusChecksConfig{
					valGen,
					&trueBool,
					nil,
				},

				&RequiredPRReviewConfig{
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
			&Config{valGen, "branch-protection-id", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
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

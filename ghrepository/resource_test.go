package ghrepository_test

import (
	"testing"

	"github.com/yoanm/go-tfsig/testutils"

	"github.com/yoanm/go-gh2tf"
	"github.com/yoanm/go-gh2tf/ghrepository"
)

func TestNew(t *testing.T) {
	t.Parallel()

	valGen := gh2tf.NewValueGenerator()
	repoName := "repository_name"
	visibility := "visibility"
	trueBool := "true"
	falseBool := "false"
	topics := []string{"topic"}
	description := "description"
	homepageUrl := "homepage-url"
	mergeCommitTitle := "merge-commit-title"
	mergeCommitMessage := "merge-commit-message"
	squashMergeCommitTitle := "squash-merge-commit-title"
	squashMergeCommitMessage := "squash-merge-commit-message"
	repoPageSourcePath := "branch"
	repoPageSourceSource := "source"
	templateOwner := "owner"
	templateRepository := "repository"
	cases := map[string]struct {
		value      *ghrepository.Config
		goldenFile string
	}{
		"Full": {
			&ghrepository.Config{
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
			},
			"full",
		},
		"Short": {
			&ghrepository.Config{
				valGen, "repository-id", &repoName, nil, nil, nil, nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
			},
			"short",
		},
		"Basic": {
			&ghrepository.Config{
				valGen,
				"repository-id",
				&repoName,
				&visibility,
				nil,
				&description,
				nil,
				&trueBool,
				nil,
				nil,
				&trueBool,
				nil,
				&topics,

				&trueBool,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				&trueBool,
				&trueBool,
				&ghrepository.PagesConfig{&ghrepository.PagesSourceConfig{valGen, nil, &repoPageSourcePath}},
				nil,
			},
			"basic",
		},
		"Enhanced": {
			&ghrepository.Config{
				valGen,
				"repository-id",
				&repoName,
				&visibility,
				nil,
				nil,
				nil,
				nil,
				&trueBool,
				&trueBool,
				nil,
				&homepageUrl,
				nil,

				nil,
				nil,
				&trueBool,
				nil,
				&trueBool,
				&mergeCommitTitle,
				nil,
				&squashMergeCommitTitle,
				nil,
				nil,
				&trueBool,
				&ghrepository.PagesConfig{
					&ghrepository.PagesSourceConfig{valGen, &repoPageSourcePath, nil},
				},
				&ghrepository.TemplateConfig{valGen, &templateOwner, &templateRepository},
			},
			"enhanced",
		},
		"Page without source": {
			&ghrepository.Config{
				valGen,
				"repository-id",
				&repoName,
				&visibility,
				nil,
				&description,
				nil,
				&trueBool,
				nil,
				nil,
				&trueBool,
				nil,
				&topics,

				&trueBool,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				&trueBool,
				&trueBool,
				&ghrepository.PagesConfig{&ghrepository.PagesSourceConfig{valGen, nil, nil}},
				nil,
			},
			"basic",
		},
		"Has not resource": {
			&ghrepository.Config{
				valGen, "repository-id", nil, nil, nil, nil, nil, nil, nil, nil,
				nil, nil,
				nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
				nil, nil, nil,
			},
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

				if err := testutils.EnsureBlockFileEqualsGoldenFile(ghrepository.New(tcase.value), tcase.goldenFile); err != nil {
					t.Errorf("Case \"%s\": %v", t.Name(), err)
				}
			},
		)
	}
}

func TestNewPagesSignature_nil(t *testing.T) {
	t.Parallel()

	got := ghrepository.NewPagesSignature(nil)
	if got != nil {
		t.Errorf("expected nil, got  %#v", got)
	}
}

func TestNewPagesSourceSignature_nil(t *testing.T) {
	t.Parallel()

	got := ghrepository.NewPagesSourceSignature(nil)
	if got != nil {
		t.Errorf("expected nil, got  %#v", got)
	}
}

func TestNewTemplateSignature_nil(t *testing.T) {
	t.Parallel()

	got := ghrepository.NewTemplateSignature(nil)
	if got != nil {
		t.Errorf("expected nil, got  %#v", got)
	}
}

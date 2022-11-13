package ghrepository

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-tfsig"
)

/** Public **/

// New returns the `github_repository` terraform resource as `hclwrite.Block`
//
// It returns `nil` if resource is empty.
func New(conf ConfigProvider) *hclwrite.Block {
	if signature := NewSignature(conf); signature != nil {
		return signature.Build()
	}

	return nil
}

// NewSignature returns the `github_branch_protection` terraform resource as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewSignature(conf ConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptyResource("github_repository", conf.ResourceIdentifier())

	repoBlock1(sig, conf)

	repoBlock2(sig, conf)

	if subSig := NewTemplateSignature(conf.TemplateConfig()); subSig != nil {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		sig.AppendChild(subSig)
	}

	repoBlock3(sig, conf)

	if subSig := NewPagesSignature(conf.PagesConfig()); subSig != nil {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		sig.AppendChild(subSig)
	}

	repoBlock4(sig, conf)

	repoBlock5(sig, conf)

	repoBlock6(sig, conf)

	if vulnerabilityAlerts := conf.VulnerabilityAlertsValue(); vulnerabilityAlerts != nil {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		sig.AppendAttribute("vulnerability_alerts", *vulnerabilityAlerts)
	}

	repoBlock7(sig, conf)

	return sig
}

// NewPagesSignature returns the `github_repository->pages` terraform block
// as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewPagesSignature(conf PagesConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptySignature("pages")

	if subSig := NewPagesSourceSignature(conf.SourceConfig()); subSig != nil {
		sig.AppendChild(subSig)
	}

	return sig
}

// NewPagesSourceSignature returns the `github_repository->pages->source` terraform block
// as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewPagesSourceSignature(conf PagesSourceConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptySignature("source")

	if v := conf.BranchValue(); v != nil {
		sig.AppendAttribute("branch", *v)
	}

	if v := conf.PathValue(); v != nil {
		sig.AppendAttribute("path", *v)
	}

	return sig
}

// NewTemplateSignature returns the `github_repository->template` terraform block
// as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewTemplateSignature(conf TemplateConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptySignature("template")

	if v := conf.OwnerValue(); v != nil {
		sig.AppendAttribute("owner", *v)
	}

	if v := conf.RepositoryValue(); v != nil {
		sig.AppendAttribute("repository", *v)
	}

	return sig
}

/** Private **/

// repoBlock1 appends `name` and `auto_init` attributes.
func repoBlock1(sig *tfsig.BlockSignature, conf ConfigProvider) {
	if name := conf.NameValue(); name != nil {
		sig.AppendAttribute("name", *name)
	}

	if autoInit := conf.AutoInitValue(); autoInit != nil {
		sig.AppendAttribute("auto_init", *autoInit)
	}
}

// repoBlock2 appends `visibility` and `description` attributes.
// It adds an empty line before in case there is existing attributes.
func repoBlock2(sig *tfsig.BlockSignature, conf ConfigProvider) {
	visibility := conf.VisibilityValue()
	description := conf.DescriptionValue()

	if visibility != nil || description != nil {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		if visibility != nil {
			sig.AppendAttribute("visibility", *visibility)
		}

		if description != nil {
			sig.AppendAttribute("description", *description)
		}
	}
}

// repoBlock3 appends `topics` and `homepage_url` attributes.
// It adds an empty line before in case there is existing attributes.
func repoBlock3(sig *tfsig.BlockSignature, conf ConfigProvider) {
	topics := conf.TopicsValue()
	homepageUrl := conf.HomepageUrlValue()

	if homepageUrl != nil || topics != nil {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		if topics != nil {
			sig.AppendAttribute("topics", *topics)
		}

		if homepageUrl != nil {
			sig.AppendAttribute("homepage_url", *homepageUrl)
		}
	}
}

// repoBlock4 appends `has_issues`, `has_projects`, `has_wiki` and `has_downloads` attributes.
// It adds an empty line before in case there is existing attributes.
func repoBlock4(sig *tfsig.BlockSignature, conf ConfigProvider) {
	hasIssues := conf.HasIssuesValue()
	hasProjects := conf.HasProjectsValue()
	hasWiki := conf.HasWikiValue()
	hasDownloads := conf.HasDownloadsValue()
	hasBlock := hasIssues != nil || hasProjects != nil || hasWiki != nil || hasDownloads != nil

	if hasBlock { //nolint:nestif // disabled as it's based on all properties in the block
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		if hasIssues != nil {
			sig.AppendAttribute("has_issues", *hasIssues)
		}

		if hasProjects != nil {
			sig.AppendAttribute("has_projects", *hasProjects)
		}

		if hasWiki != nil {
			sig.AppendAttribute("has_wiki", *hasWiki)
		}

		if hasDownloads != nil {
			sig.AppendAttribute("has_downloads", *hasDownloads)
		}
	}
}

// repoBlock5 appends `allow_merge_commit`, `allow_rebase_merge`, `allow_squash_merge`, `allow_auto_merge`
// and `delete_branch_on_merge` attributes.
// It adds an empty line before in case there is existing attributes.
//
//nolint:gocognit,cyclop // disabled as complexity is mostly based on number of terraform attributes
func repoBlock5(sig *tfsig.BlockSignature, conf ConfigProvider) {
	allowMergeCommit := conf.AllowMergeCommitValue()
	allowRebaseMerge := conf.AllowRebaseMergeValue()
	allowSquashMerge := conf.AllowSquashMergeValue()
	allowAutoMerge := conf.AllowAutoMergeValue()
	deleteBranchOnMerge := conf.DeleteBranchOnMergeValue()
	hasBlock := allowMergeCommit != nil || allowRebaseMerge != nil || allowSquashMerge != nil ||
		allowAutoMerge != nil || deleteBranchOnMerge != nil

	if hasBlock { //nolint:nestif // disabled as it's based on all properties in the block
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		if allowMergeCommit != nil {
			sig.AppendAttribute("allow_merge_commit", *allowMergeCommit)
		}

		if allowRebaseMerge != nil {
			sig.AppendAttribute("allow_rebase_merge", *allowRebaseMerge)
		}

		if allowSquashMerge != nil {
			sig.AppendAttribute("allow_squash_merge", *allowSquashMerge)
		}

		if allowAutoMerge != nil {
			sig.AppendAttribute("allow_auto_merge", *allowAutoMerge)
		}

		if deleteBranchOnMerge != nil {
			sig.AppendAttribute("delete_branch_on_merge", *deleteBranchOnMerge)
		}
	}
}

// repoBlock6 appends `merge_commit_title`, `merge_commit_message`, `squash_merge_commit_title`
// and `squash_merge_commit_message` attributes.
// It adds an empty line before in case there is existing attributes.
func repoBlock6(sig *tfsig.BlockSignature, conf ConfigProvider) {
	mergeCommitTitle := conf.MergeCommitTitleValue()
	mergeCommitMessage := conf.MergeCommitMessageValue()
	squashMergeCommitTitle := conf.SquashMergeCommitTitleValue()
	squashMergeCommitMessage := conf.SquashMergeCommitMessageValue()
	hasBlock := mergeCommitTitle != nil || mergeCommitMessage != nil || squashMergeCommitTitle != nil ||
		squashMergeCommitMessage != nil

	if hasBlock { //nolint:nestif // disabled as it's based on all properties in the block
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		if mergeCommitTitle != nil {
			sig.AppendAttribute("merge_commit_title", *mergeCommitTitle)
		}

		if mergeCommitMessage != nil {
			sig.AppendAttribute("merge_commit_message", *mergeCommitMessage)
		}

		if squashMergeCommitTitle != nil {
			sig.AppendAttribute("squash_merge_commit_title", *squashMergeCommitTitle)
		}

		if squashMergeCommitMessage != nil {
			sig.AppendAttribute("squash_merge_commit_message", *squashMergeCommitMessage)
		}
	}
}

// repoBlock7 appends `archived` and `archive_on_destroy` attributes.
// It adds an empty line before in case there is existing attributes.
func repoBlock7(sig *tfsig.BlockSignature, conf ConfigProvider) {
	archived := conf.ArchivedValue()
	archiveOnDestroy := conf.ArchiveOnDestroyValue()

	if archived != nil || archiveOnDestroy != nil {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		if archived != nil {
			sig.AppendAttribute("archived", *archived)
		}

		if archiveOnDestroy != nil {
			sig.AppendAttribute("archive_on_destroy", *archiveOnDestroy)
		}
	}
}

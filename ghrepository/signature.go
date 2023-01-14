package ghrepository

import (
	"github.com/yoanm/go-tfsig"
)

/** Public **/

// NewSignature returns the `github_branch_protection` terraform resource as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewSignature(conf ConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewResource("github_repository", conf.ResourceIdentifier())

	repoBlock1(sig, conf)

	repoBlock2(sig, conf)

	tfsig.AppendChildIfNotNil(sig, NewTemplateSignature(conf.TemplateConfig()))

	repoBlock3(sig, conf)

	tfsig.AppendChildIfNotNil(sig, NewPagesSignature(conf.PagesConfig()))

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

	sig := tfsig.NewSignature("pages")

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

	sig := tfsig.NewSignature("source")

	tfsig.AppendAttributeIfNotNil(sig, "branch", conf.BranchValue())
	tfsig.AppendAttributeIfNotNil(sig, "path", conf.PathValue())

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

	sig := tfsig.NewSignature("template")

	tfsig.AppendAttributeIfNotNil(sig, "owner", conf.OwnerValue())
	tfsig.AppendAttributeIfNotNil(sig, "repository", conf.RepositoryValue())

	return sig
}

/** Private **/

// repoBlock1 appends `name` and `auto_init` attributes.
func repoBlock1(sig *tfsig.BlockSignature, conf ConfigProvider) {
	tfsig.AppendAttributeIfNotNil(sig, "name", conf.NameValue())
	tfsig.AppendAttributeIfNotNil(sig, "auto_init", conf.AutoInitValue())
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

		tfsig.AppendAttributeIfNotNil(sig, "visibility", visibility)
		tfsig.AppendAttributeIfNotNil(sig, "description", description)
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

		tfsig.AppendAttributeIfNotNil(sig, "topics", topics)
		tfsig.AppendAttributeIfNotNil(sig, "homepage_url", homepageUrl)
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

	if hasBlock {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		tfsig.AppendAttributeIfNotNil(sig, "has_issues", hasIssues)
		tfsig.AppendAttributeIfNotNil(sig, "has_projects", hasProjects)
		tfsig.AppendAttributeIfNotNil(sig, "has_wiki", hasWiki)
		tfsig.AppendAttributeIfNotNil(sig, "has_downloads", hasDownloads)
	}
}

// repoBlock5 appends `allow_merge_commit`, `allow_rebase_merge`, `allow_squash_merge`, `allow_auto_merge`
// and `delete_branch_on_merge` attributes.
// It adds an empty line before in case there is existing attributes.
func repoBlock5(sig *tfsig.BlockSignature, conf ConfigProvider) {
	allowMergeCommit := conf.AllowMergeCommitValue()
	allowRebaseMerge := conf.AllowRebaseMergeValue()
	allowSquashMerge := conf.AllowSquashMergeValue()
	allowAutoMerge := conf.AllowAutoMergeValue()
	deleteBranchOnMerge := conf.DeleteBranchOnMergeValue()
	hasBlock := allowMergeCommit != nil || allowRebaseMerge != nil || allowSquashMerge != nil ||
		allowAutoMerge != nil || deleteBranchOnMerge != nil

	if hasBlock {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		tfsig.AppendAttributeIfNotNil(sig, "allow_merge_commit", allowMergeCommit)
		tfsig.AppendAttributeIfNotNil(sig, "allow_rebase_merge", allowRebaseMerge)
		tfsig.AppendAttributeIfNotNil(sig, "allow_squash_merge", allowSquashMerge)
		tfsig.AppendAttributeIfNotNil(sig, "allow_auto_merge", allowAutoMerge)
		tfsig.AppendAttributeIfNotNil(sig, "delete_branch_on_merge", deleteBranchOnMerge)
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

	if hasBlock {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		tfsig.AppendAttributeIfNotNil(sig, "merge_commit_title", mergeCommitTitle)
		tfsig.AppendAttributeIfNotNil(sig, "merge_commit_message", mergeCommitMessage)
		tfsig.AppendAttributeIfNotNil(sig, "squash_merge_commit_title", squashMergeCommitTitle)
		tfsig.AppendAttributeIfNotNil(sig, "squash_merge_commit_message", squashMergeCommitMessage)
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

		tfsig.AppendAttributeIfNotNil(sig, "archived", archived)
		tfsig.AppendAttributeIfNotNil(sig, "archive_on_destroy", archiveOnDestroy)
	}
}

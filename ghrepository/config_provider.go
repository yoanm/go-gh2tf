package ghrepository

import "github.com/zclconf/go-cty/cty"

// ConfigProvider defines required methods to be used when creating `github_repository` terraform resource.
//
//nolint:interfacebloat // Disabled as number of methods is based on number of terraform attributes
type ConfigProvider interface {
	// HasResource will indicate if resource has to be created (return `true` for instance to always create the
	// resource even if empty)
	HasResource() bool
	// ResourceIdentifier returns the resource identifier
	ResourceIdentifier() string

	// NameValue return the `github_repository` `name` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	NameValue() *cty.Value
	// VisibilityValue return the `github_repository` `visbility` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	VisibilityValue() *cty.Value
	// ArchivedValue return the `github_repository` `archived` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	ArchivedValue() *cty.Value
	// DescriptionValue return the `github_repository` `description` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	DescriptionValue() *cty.Value
	// AutoInitValue return the `github_repository` `auto_init` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	AutoInitValue() *cty.Value
	// HasIssuesValue return the `github_repository` `has_issues` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	HasIssuesValue() *cty.Value
	// HasProjectsValue return the `github_repository` `has_projects` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	HasProjectsValue() *cty.Value
	// HasWikiValue return the `github_repository` `has_wiki` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	HasWikiValue() *cty.Value
	// HasDownloadsValue return the `github_repository` `has_downloads` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	HasDownloadsValue() *cty.Value
	// HomepageUrlValue return the `github_repository` `homepage_url` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	HomepageUrlValue() *cty.Value
	// TopicsValue return the `github_repository` `topics` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	TopicsValue() *cty.Value
	// VulnerabilityAlertsValue return the `github_repository` `vulnerability` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	VulnerabilityAlertsValue() *cty.Value
	// AllowMergeCommitValue return the `github_repository` `allow_merge_commit` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	AllowMergeCommitValue() *cty.Value
	// AllowRebaseMergeValue return the `github_repository` `allow_rebase_merge` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	AllowRebaseMergeValue() *cty.Value
	// AllowSquashMergeValue return the `github_repository` `allow_squash_merge` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	AllowSquashMergeValue() *cty.Value
	// AllowAutoMergeValue return the `github_repository` `allow_auto_merge` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	AllowAutoMergeValue() *cty.Value
	// MergeCommitTitleValue return the `github_repository` `merge_commit_title` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	MergeCommitTitleValue() *cty.Value
	// MergeCommitMessageValue return the `github_repository` `merge_commit_message` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	MergeCommitMessageValue() *cty.Value
	// SquashMergeCommitTitleValue return the `github_repository` `squash_merge_commit_title` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	SquashMergeCommitTitleValue() *cty.Value
	// SquashMergeCommitMessageValue return the `github_repository` `squash_merge_commit_message` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	SquashMergeCommitMessageValue() *cty.Value
	// DeleteBranchOnMergeValue return the `github_repository` `delete_branch_on_merge` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	DeleteBranchOnMergeValue() *cty.Value
	// ArchiveOnDestroyValue return the `github_repository` `archive_on_destroy` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	ArchiveOnDestroyValue() *cty.Value

	// PagesConfig return the config provider for `github_repository->pages` block
	PagesConfig() PagesConfigProvider
	// TemplateConfig return the config provider for `github_repository->template` block
	TemplateConfig() TemplateConfigProvider
}

// PagesConfigProvider defines required methods to be used when creating `github_repository->pages` terraform block.
type PagesConfigProvider interface {
	// HasResource will indicate if resource has to be created (return `true` for instance to always create the
	// resource even if empty)
	HasResource() bool

	// SourceConfig return the config provider for `github_repository->pages->source` block
	SourceConfig() PagesSourceConfigProvider
}

// PagesSourceConfigProvider defines required methods to be used when creating `github_repository->pages->source`
// terraform block.
type PagesSourceConfigProvider interface {
	// HasResource will indicate if resource has to be created (return `true` for instance to always create the
	// resource even if empty)
	HasResource() bool

	// BranchValue return the `github_repository->pages->source` `branch` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	BranchValue() *cty.Value

	// PathValue return the `github_repository->pages->source` `path` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	PathValue() *cty.Value
}

// TemplateConfigProvider defines required methods to be used when creating `github_repository->template`
// terraform block.
type TemplateConfigProvider interface {
	// HasResource will indicate if resource has to be created (return `true` for instance to always create the
	// resource even if empty)
	HasResource() bool

	// OwnerValue return the `github_repository->template` `owner` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	OwnerValue() *cty.Value

	// RepositoryValue return the `github_repository->template` `repository` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	RepositoryValue() *cty.Value
}

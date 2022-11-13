/*
Package ghrepository provides methods to create `github_repository` terraform resource
*/
package ghrepository

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig"
)

// Config is the default implementation of `ConfigProvider`.
type Config struct {
	ValueGenerator tfsig.ValueGenerator
	Identifier     string

	Name                     *string
	Visibility               *string
	Archived                 *string
	Description              *string
	AutoInit                 *string
	HasIssues                *string
	HasProjects              *string
	HasWiki                  *string
	HasDownloads             *string
	HomepageUrl              *string
	Topics                   *[]string
	VulnerabilityAlerts      *string
	AllowMergeCommit         *string
	AllowRebaseMerge         *string
	AllowSquashMerge         *string
	AllowAutoMerge           *string
	MergeCommitTitle         *string
	MergeCommitMessage       *string
	SquashMergeCommitTitle   *string
	SquashMergeCommitMessage *string
	DeleteBranchOnMerge      *string
	ArchiveOnDestroy         *string
	Page                     *PagesConfig
	Template                 *TemplateConfig
}

// PagesConfig is the default implementation of `PagesConfigProvider`.
type PagesConfig struct {
	Source *PagesSourceConfig
}

// PagesSourceConfig is the default implementation of `PagesSourceConfigProvider`.
type PagesSourceConfig struct {
	ValueGenerator tfsig.ValueGenerator
	Branch         *string
	Path           *string
}

// TemplateConfig is the default implementation of `TemplateConfigProvider`.
type TemplateConfig struct {
	ValueGenerator tfsig.ValueGenerator
	Owner          *string
	Repository     *string
}

// HasResource returns `true` in case at least `Name` value exists, else `false`.
func (c *Config) HasResource() bool {
	// Repository being the only required value
	return c != nil && c.Name != nil
}

// ResourceIdentifier returns the provided terraform resource identifier.
func (c *Config) ResourceIdentifier() string {
	return c.Identifier
}

// NameValue returns the provided `github_repository` `name` attribute value as `cty.String`
// or `nil` if not provided.
func (c *Config) NameValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Name)
}

// VisibilityValue returns the provided `github_repository` `visibility` attribute value as `cty.String`
// or `nil` if not provided.
func (c *Config) VisibilityValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Visibility)
}

// ArchivedValue returns the provided `github_repository` `archived` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) ArchivedValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.Archived)
}

// DescriptionValue returns the provided `github_repository` `description` attribute value as `cty.String`
// or `nil` if not provided.
func (c *Config) DescriptionValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Description)
}

// AutoInitValue returns the provided `github_repository` `auto_init` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) AutoInitValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.AutoInit)
}

// HasIssuesValue returns the provided `github_repository` `has_issues` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) HasIssuesValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.HasIssues)
}

// HasProjectsValue returns the provided `github_repository` `has_projects` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) HasProjectsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.HasProjects)
}

// HasWikiValue returns the provided `github_repository` `has_wiki` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) HasWikiValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.HasWiki)
}

// HasDownloadsValue returns the provided `github_repository` `has_downloads` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) HasDownloadsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.HasDownloads)
}

// HomepageUrlValue returns the provided `github_repository` `homepage_url` attribute value as `cty.String`
// or `nil` if not provided.
func (c *Config) HomepageUrlValue() *cty.Value {
	return c.ValueGenerator.ToString(c.HomepageUrl)
}

// TopicsValue returns the provided `github_repository` `topics` attribute value
// as `cty.List` of `cty.String` or `nil` if not provided.
func (c *Config) TopicsValue() *cty.Value {
	return c.ValueGenerator.ToStringList(c.Topics)
}

// PagesConfig returns the provided `PagesConfigProvider`.
func (c *Config) PagesConfig() PagesConfigProvider { //nolint:ireturn,lll // Disabled as value is passed to a method expecting that interface
	return c.Page
}

// VulnerabilityAlertsValue returns the provided `github_repository` `vulnerabililty_alerts` attribute value
// as `cty.Bool` or `nil` if not provided.
func (c *Config) VulnerabilityAlertsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.VulnerabilityAlerts)
}

// AllowMergeCommitValue returns the provided `github_repository` `allow_merge_commit` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) AllowMergeCommitValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.AllowMergeCommit)
}

// AllowRebaseMergeValue returns the provided `github_repository` `allow_rebase_merge` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) AllowRebaseMergeValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.AllowRebaseMerge)
}

// AllowSquashMergeValue returns the provided `github_repository` `allow_squash_merge` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) AllowSquashMergeValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.AllowSquashMerge)
}

// AllowAutoMergeValue returns the provided `github_repository` `allow_auto_merge` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) AllowAutoMergeValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.AllowAutoMerge)
}

// MergeCommitTitleValue returns the provided `github_repository` `merge_commit_title` attribute value as `cty.String`
// or `nil` if not provided.
func (c *Config) MergeCommitTitleValue() *cty.Value {
	return c.ValueGenerator.ToString(c.MergeCommitTitle)
}

// MergeCommitMessageValue returns the provided `github_repository` `merge_commit_message` attribute value
// as `cty.String` or `nil` if not provided.
func (c *Config) MergeCommitMessageValue() *cty.Value {
	return c.ValueGenerator.ToString(c.MergeCommitMessage)
}

// SquashMergeCommitTitleValue returns the provided `github_repository` `squash_merge_commit_title` attribute value
// as `cty.String` or `nil` if not provided.
func (c *Config) SquashMergeCommitTitleValue() *cty.Value {
	return c.ValueGenerator.ToString(c.SquashMergeCommitTitle)
}

// SquashMergeCommitMessageValue returns the provided `github_repository` `squash_merge_commit_message` attribute value
// as `cty.String` or `nil` if not provided.
func (c *Config) SquashMergeCommitMessageValue() *cty.Value {
	return c.ValueGenerator.ToString(c.SquashMergeCommitMessage)
}

// DeleteBranchOnMergeValue returns the provided `github_repository` `delete_branch_on_merge` attribute value
// as `cty.Bool` or `nil` if not provided.
func (c *Config) DeleteBranchOnMergeValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.DeleteBranchOnMerge)
}

// TemplateConfig returns the provided `TemplateConfigProvider`.
func (c *Config) TemplateConfig() TemplateConfigProvider { //nolint:ireturn,lll // Disabled as value is passed to a method expecting that interface
	return c.Template
}

// ArchiveOnDestroyValue returns the provided `github_repository` `archive_on_destroy` attribute value as `cty.Bool`
// or `nil` if not provided.
func (c *Config) ArchiveOnDestroyValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.ArchiveOnDestroy)
}

// PagesConfigProvider implementation

// HasResource returns `true` in case `Source` config exists and has resource, else `false`.
func (c *PagesConfig) HasResource() bool {
	return c != nil && c.Source != nil && c.Source.HasResource()
}

// SourceConfig returns the provided `PagesSourceConfigProvider`.
func (c *PagesConfig) SourceConfig() PagesSourceConfigProvider { //nolint:ireturn,lll // Disabled as value is passed to a method expecting that interface
	return c.Source
}

// PagesSourceConfigProvider implementation

// HasResource returns `true` in case `Branch` value exists, else `false`.
func (c *PagesSourceConfig) HasResource() bool {
	return c != nil && c.Branch != nil
}

// BranchValue returns the provided `github_repository->pages->source` `branch` attribute value
// as `cty.String` or `nil` if not provided.
func (c *PagesSourceConfig) BranchValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Branch)
}

// PathValue returns the provided `github_repository->pages->source` `path` attribute value
// as `cty.String` or `nil` if not provided.
func (c *PagesSourceConfig) PathValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Path)
}

// TemplateConfigProvider implementation

// HasResource returns `true` in case both `Repository` and `Owner` values exist, else `false`.
func (c *TemplateConfig) HasResource() bool {
	return c != nil && c.Repository != nil && c.Owner != nil
}

// OwnerValue returns the provided `github_repository->template` `owner` attribute value as `cty.String`
// or `nil` if not provided.
func (c *TemplateConfig) OwnerValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Owner)
}

// RepositoryValue returns the provided `github_repository->template` `repository` attribute value as `cty.String`
// or `nil` if not provided.
func (c *TemplateConfig) RepositoryValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Repository)
}

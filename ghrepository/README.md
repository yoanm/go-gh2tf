# ghrepository

Package ghrepository provides methods to create `github_repository` terraform resource

## Functions

### func [New](./resource.go#L13)

`func New(conf ConfigProvider) *hclwrite.Block`

New returns the `github_repository` terraform resource as `hclwrite.Block`

It returns `nil` if resource is empty.

```golang
valGen := gh2tf.NewValueGenerator()
repoName := "repository_name"
visibility := "var.default_visibility"
trueBool := "true"
falseBool := "false"
topics := []string{"topic"}
description := "local.my_local"
homepageUrl := "homepage-url"
mergeCommitTitle := "merge-commit-title"
mergeCommitMessage := "merge-commit-message"
squashMergeCommitTitle := "squash-merge-commit-title"
squashMergeCommitMessage := "squash-merge-commit-message"
repoPageSourcePath := "branch"
repoPageSourceSource := "source"
templateOwner := "owner"
templateRepository := "repository"
res := &ghrepository.Config{
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
}

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(ghrepository.New(res))
fmt.Println(string(hclFile.Bytes()))
```

 Output:

```terraform
resource "github_repository" "repository-id" {
  name      = "repository_name"
  auto_init = false

  visibility  = var.default_visibility
  description = local.my_local

  template {
    owner      = "owner"
    repository = "repository"
  }

  topics       = ["topic"]
  homepage_url = "homepage-url"

  pages {
    source {
      branch = "branch"
      path   = "source"
    }
  }

  has_issues    = true
  has_projects  = false
  has_wiki      = true
  has_downloads = false

  allow_merge_commit     = false
  allow_rebase_merge     = true
  allow_squash_merge     = false
  allow_auto_merge       = true
  delete_branch_on_merge = false

  merge_commit_title          = "merge-commit-title"
  merge_commit_message        = "merge-commit-message"
  squash_merge_commit_title   = "squash-merge-commit-title"
  squash_merge_commit_message = "squash-merge-commit-message"

  vulnerability_alerts = true

  archived           = false
  archive_on_destroy = true
}
```

### func [NewPagesSignature](./resource.go#L76)

`func NewPagesSignature(conf PagesConfigProvider) *tfsig.BlockSignature`

NewPagesSignature returns the `github_repository->pages` terraform block
as `tfsig.BlockSignature`

It returns `nil` if resource is empty.

### func [NewPagesSourceSignature](./resource.go#L94)

`func NewPagesSourceSignature(conf PagesSourceConfigProvider) *tfsig.BlockSignature`

NewPagesSourceSignature returns the `github_repository->pages->source` terraform block
as `tfsig.BlockSignature`

It returns `nil` if resource is empty.

### func [NewSignature](./resource.go#L24)

`func NewSignature(conf ConfigProvider) *tfsig.BlockSignature`

NewSignature returns the `github_branch_protection` terraform resource as `tfsig.BlockSignature`

It returns `nil` if resource is empty.

### func [NewTemplateSignature](./resource.go#L116)

`func NewTemplateSignature(conf TemplateConfigProvider) *tfsig.BlockSignature`

NewTemplateSignature returns the `github_repository->template` terraform block
as `tfsig.BlockSignature`

It returns `nil` if resource is empty.

## Types

### type [Config](./config.go#L12)

`type Config struct { ... }`

Config is the default implementation of `ConfigProvider`.

#### func (*Config) [AllowAutoMergeValue](./config.go#L169)

`func (c *Config) AllowAutoMergeValue() *cty.Value`

AllowAutoMergeValue returns the provided `github_repository` `allow_auto_merge` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [AllowMergeCommitValue](./config.go#L151)

`func (c *Config) AllowMergeCommitValue() *cty.Value`

AllowMergeCommitValue returns the provided `github_repository` `allow_merge_commit` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [AllowRebaseMergeValue](./config.go#L157)

`func (c *Config) AllowRebaseMergeValue() *cty.Value`

AllowRebaseMergeValue returns the provided `github_repository` `allow_rebase_merge` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [AllowSquashMergeValue](./config.go#L163)

`func (c *Config) AllowSquashMergeValue() *cty.Value`

AllowSquashMergeValue returns the provided `github_repository` `allow_squash_merge` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [ArchiveOnDestroyValue](./config.go#L210)

`func (c *Config) ArchiveOnDestroyValue() *cty.Value`

ArchiveOnDestroyValue returns the provided `github_repository` `archive_on_destroy` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [ArchivedValue](./config.go#L86)

`func (c *Config) ArchivedValue() *cty.Value`

ArchivedValue returns the provided `github_repository` `archived` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [AutoInitValue](./config.go#L98)

`func (c *Config) AutoInitValue() *cty.Value`

AutoInitValue returns the provided `github_repository` `auto_init` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [DeleteBranchOnMergeValue](./config.go#L199)

`func (c *Config) DeleteBranchOnMergeValue() *cty.Value`

DeleteBranchOnMergeValue returns the provided `github_repository` `delete_branch_on_merge` attribute value
as `cty.Bool` or `nil` if not provided.

#### func (*Config) [DescriptionValue](./config.go#L92)

`func (c *Config) DescriptionValue() *cty.Value`

DescriptionValue returns the provided `github_repository` `description` attribute value as `cty.String`
or `nil` if not provided.

#### func (*Config) [HasDownloadsValue](./config.go#L122)

`func (c *Config) HasDownloadsValue() *cty.Value`

HasDownloadsValue returns the provided `github_repository` `has_downloads` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [HasIssuesValue](./config.go#L104)

`func (c *Config) HasIssuesValue() *cty.Value`

HasIssuesValue returns the provided `github_repository` `has_issues` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [HasProjectsValue](./config.go#L110)

`func (c *Config) HasProjectsValue() *cty.Value`

HasProjectsValue returns the provided `github_repository` `has_projects` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [HasResource](./config.go#L62)

`func (c *Config) HasResource() bool`

HasResource returns `true` in case at least `Name` value exists, else `false`.

#### func (*Config) [HasWikiValue](./config.go#L116)

`func (c *Config) HasWikiValue() *cty.Value`

HasWikiValue returns the provided `github_repository` `has_wiki` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [HomepageUrlValue](./config.go#L128)

`func (c *Config) HomepageUrlValue() *cty.Value`

HomepageUrlValue returns the provided `github_repository` `homepage_url` attribute value as `cty.String`
or `nil` if not provided.

#### func (*Config) [MergeCommitMessageValue](./config.go#L181)

`func (c *Config) MergeCommitMessageValue() *cty.Value`

MergeCommitMessageValue returns the provided `github_repository` `merge_commit_message` attribute value
as `cty.String` or `nil` if not provided.

#### func (*Config) [MergeCommitTitleValue](./config.go#L175)

`func (c *Config) MergeCommitTitleValue() *cty.Value`

MergeCommitTitleValue returns the provided `github_repository` `merge_commit_title` attribute value as `cty.String`
or `nil` if not provided.

#### func (*Config) [NameValue](./config.go#L74)

`func (c *Config) NameValue() *cty.Value`

NameValue returns the provided `github_repository` `name` attribute value as `cty.String`
or `nil` if not provided.

#### func (*Config) [PagesConfig](./config.go#L139)

`func (c *Config) PagesConfig() PagesConfigProvider`

PagesConfig returns the provided `PagesConfigProvider`.

#### func (*Config) [ResourceIdentifier](./config.go#L68)

`func (c *Config) ResourceIdentifier() string`

ResourceIdentifier returns the provided terraform resource identifier.

#### func (*Config) [SquashMergeCommitMessageValue](./config.go#L193)

`func (c *Config) SquashMergeCommitMessageValue() *cty.Value`

SquashMergeCommitMessageValue returns the provided `github_repository` `squash_merge_commit_message` attribute value
as `cty.String` or `nil` if not provided.

#### func (*Config) [SquashMergeCommitTitleValue](./config.go#L187)

`func (c *Config) SquashMergeCommitTitleValue() *cty.Value`

SquashMergeCommitTitleValue returns the provided `github_repository` `squash_merge_commit_title` attribute value
as `cty.String` or `nil` if not provided.

#### func (*Config) [TemplateConfig](./config.go#L204)

`func (c *Config) TemplateConfig() TemplateConfigProvider`

TemplateConfig returns the provided `TemplateConfigProvider`.

#### func (*Config) [TopicsValue](./config.go#L134)

`func (c *Config) TopicsValue() *cty.Value`

TopicsValue returns the provided `github_repository` `topics` attribute value
as `cty.List` of `cty.String` or `nil` if not provided.

#### func (*Config) [VisibilityValue](./config.go#L80)

`func (c *Config) VisibilityValue() *cty.Value`

VisibilityValue returns the provided `github_repository` `visibility` attribute value as `cty.String`
or `nil` if not provided.

#### func (*Config) [VulnerabilityAlertsValue](./config.go#L145)

`func (c *Config) VulnerabilityAlertsValue() *cty.Value`

VulnerabilityAlertsValue returns the provided `github_repository` `vulnerabililty_alerts` attribute value
as `cty.Bool` or `nil` if not provided.

### type [ConfigProvider](./config_provider.go#L8)

`type ConfigProvider interface { ... }`

ConfigProvider defines required methods to be used when creating `github_repository` terraform resource.

### type [PagesConfig](./config.go#L43)

`type PagesConfig struct { ... }`

PagesConfig is the default implementation of `PagesConfigProvider`.

#### func (*PagesConfig) [HasResource](./config.go#L217)

`func (c *PagesConfig) HasResource() bool`

HasResource returns `true` in case `Source` config exists and has resource, else `false`.

#### func (*PagesConfig) [SourceConfig](./config.go#L222)

`func (c *PagesConfig) SourceConfig() PagesSourceConfigProvider`

SourceConfig returns the provided `PagesSourceConfigProvider`.

### type [PagesConfigProvider](./config_provider.go#L111)

`type PagesConfigProvider interface { ... }`

PagesConfigProvider defines required methods to be used when creating `github_repository->pages` terraform block.

### type [PagesSourceConfig](./config.go#L48)

`type PagesSourceConfig struct { ... }`

PagesSourceConfig is the default implementation of `PagesSourceConfigProvider`.

#### func (*PagesSourceConfig) [BranchValue](./config.go#L235)

`func (c *PagesSourceConfig) BranchValue() *cty.Value`

BranchValue returns the provided `github_repository->pages->source` `branch` attribute value
as `cty.String` or `nil` if not provided.

#### func (*PagesSourceConfig) [HasResource](./config.go#L229)

`func (c *PagesSourceConfig) HasResource() bool`

HasResource returns `true` in case `Branch` value exists, else `false`.

#### func (*PagesSourceConfig) [PathValue](./config.go#L241)

`func (c *PagesSourceConfig) PathValue() *cty.Value`

PathValue returns the provided `github_repository->pages->source` `path` attribute value
as `cty.String` or `nil` if not provided.

### type [PagesSourceConfigProvider](./config_provider.go#L122)

`type PagesSourceConfigProvider interface { ... }`

PagesSourceConfigProvider defines required methods to be used when creating `github_repository->pages->source`
terraform block.

### type [TemplateConfig](./config.go#L55)

`type TemplateConfig struct { ... }`

TemplateConfig is the default implementation of `TemplateConfigProvider`.

#### func (*TemplateConfig) [HasResource](./config.go#L248)

`func (c *TemplateConfig) HasResource() bool`

HasResource returns `true` in case both `Repository` and `Owner` values exist, else `false`.

#### func (*TemplateConfig) [OwnerValue](./config.go#L254)

`func (c *TemplateConfig) OwnerValue() *cty.Value`

OwnerValue returns the provided `github_repository->template` `owner` attribute value as `cty.String`
or `nil` if not provided.

#### func (*TemplateConfig) [RepositoryValue](./config.go#L260)

`func (c *TemplateConfig) RepositoryValue() *cty.Value`

RepositoryValue returns the provided `github_repository->template` `repository` attribute value as `cty.String`
or `nil` if not provided.

### type [TemplateConfigProvider](./config_provider.go#L140)

`type TemplateConfigProvider interface { ... }`

TemplateConfigProvider defines required methods to be used when creating `github_repository->template`
terraform block.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)

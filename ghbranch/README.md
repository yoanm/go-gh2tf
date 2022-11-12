# ghbranch

Package ghbranch provides methods to create `github_branch` terraform resource

## Functions

### func [New](./resource.go#L11)

`func New(c ConfigProvider) *hclwrite.Block`

New returns the `github_branch` terraform resource as `hclwrite.Block`

It returns `nil` if `NewSignature` returns nil

```golang
valGen := gh2tf.NewValueGenerator()
repoName := "repository_name"
branchName := "my_default_branch_name"
sourceBranch := "source_branch_name"
sourceSha := "a_sha"
repoNameAsLink := "github_repository.res-id.name"
branchNameAsLink := "github_branch.res-id.branch"

res := &Config{
    valGen,
    "res-id",
    &repoName,
    &branchName,
    nil,
    nil,
}
resWithSourceBranch := &Config{
    valGen,
    "res-with-source-branch",
    &repoNameAsLink,
    &branchName,
    &sourceBranch,
    nil,
}
resWithSourceSha := &Config{
    valGen,
    "res-with-source-sha",
    &repoName,
    &branchNameAsLink,
    nil,
    &sourceSha,
}

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(New(res))
hclFile.Body().AppendBlock(New(resWithSourceBranch))
hclFile.Body().AppendBlock(New(resWithSourceSha))
fmt.Println(string(hclFile.Bytes()))
```

 Output:

```terraform
resource "github_branch" "res-id" {
  repository = "repository_name"
  branch     = "my_default_branch_name"
}
resource "github_branch" "res-with-source-branch" {
  repository    = github_repository.res-id.name
  branch        = "my_default_branch_name"
  source_branch = "source_branch_name"
}
resource "github_branch" "res-with-source-sha" {
  repository = "repository_name"
  branch     = github_branch.res-id.branch
  source_sha = "a_sha"
}
```

### func [NewSignature](./resource.go#L22)

`func NewSignature(c ConfigProvider) *tfsig.BlockSignature`

NewSignature returns the `github_branch` terraform resource as `tfsig.BlockSignature`

It returns `nil` if resource is empty

## Types

### type [Config](./config.go#L12)

`type Config struct { ... }`

Config is the default implementation of `ConfigProvider`

#### func (*Config) [BranchValue](./config.go#L37)

`func (c *Config) BranchValue() *cty.Value`

BranchValue return the provided `github_branch` `branch` attribute value as `cty.String` or `nil` if not provided

#### func (*Config) [HasResource](./config.go#L22)

`func (c *Config) HasResource() bool`

HasResource returns `true` in case at least `Config` and `Repository` value exist, else `false`

#### func (*Config) [RepositoryValue](./config.go#L32)

`func (c *Config) RepositoryValue() *cty.Value`

RepositoryValue return the provided `github_branch` `repository` attribute value as `cty.String` or `nil` if not provided

#### func (*Config) [ResourceIdentifier](./config.go#L27)

`func (c *Config) ResourceIdentifier() string`

ResourceIdentifier returns the provided terraform resource identifier

#### func (*Config) [SourceBranchValue](./config.go#L42)

`func (c *Config) SourceBranchValue() *cty.Value`

SourceBranchValue return the provided `github_branch` `source_branch` attribute value as `cty.String` or `nil` if not provided

#### func (*Config) [SourceShaValue](./config.go#L47)

`func (c *Config) SourceShaValue() *cty.Value`

SourceShaValue return the provided `github_branch` `source_sha` attribute value as `cty.String` or `nil` if not provided

### type [ConfigProvider](./config_provider.go#L6)

`type ConfigProvider interface { ... }`

ConfigProvider defines required methods to be used when creating `github_branch` terraform resource

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
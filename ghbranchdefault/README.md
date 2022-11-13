# ghbranchdefault

Package ghbranchdefault provides methods to create `github_branch_default` terraform resource

## Functions

### func [New](./resource.go#L10)

`func New(c ConfigProvider) *hclwrite.Block`

New returns the `github_branch_default` terraform resource as `hclwrite.Block`

It returns `nil` if resource is empty.

```golang
valGen := gh2tf.NewValueGenerator()
repoName := "repository_name"
branchName := "my_default_branch_name"

res := &ghbranchdefault.Config{
    valGen,
    "res-id",
    &repoName,
    &branchName,
}
repoNameAsLink := "github_repository.res-id.name"
branchNameAsLink := "github_branch.res-id.branch"

resWithLinks := &ghbranchdefault.Config{
    valGen,
    "res-id-with-links",
    &repoNameAsLink,
    &branchNameAsLink,
}

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(ghbranchdefault.New(res))
hclFile.Body().AppendBlock(ghbranchdefault.New(resWithLinks))
fmt.Println(string(hclFile.Bytes()))
```

 Output:

```terraform
resource "github_branch_default" "res-id" {
  repository = "repository_name"
  branch     = "my_default_branch_name"
}
resource "github_branch_default" "res-id-with-links" {
  repository = github_repository.res-id.name
  branch     = github_branch.res-id.branch
}
```

### func [NewSignature](./signature.go#L10)

`func NewSignature(conf ConfigProvider) *tfsig.BlockSignature`

NewSignature returns the `github_branch_default` terraform resource as `tfsig.BlockSignature`

It returns `nil` if resource is empty.

## Types

### type [Config](./config.go#L12)

`type Config struct { ... }`

Config is the default implementation of `ConfigProvider`.

#### func (*Config) [BranchValue](./config.go#L37)

`func (c *Config) BranchValue() *cty.Value`

BranchValue return the provided `github_branch_default` `branch` attribute value as `cty.String`
or `nil` if not provided.

#### func (*Config) [HasResource](./config.go#L20)

`func (c *Config) HasResource() bool`

HasResource returns `true` in case both `BranchConfig` and `Repository` value exist, else `false`.

#### func (*Config) [RepositoryValue](./config.go#L31)

`func (c *Config) RepositoryValue() *cty.Value`

RepositoryValue return the provided `github_branch_default` `repository` attribute value as `cty.String`
or `nil` if not provided.

#### func (*Config) [ResourceIdentifier](./config.go#L25)

`func (c *Config) ResourceIdentifier() string`

ResourceIdentifier returns the provided terraform resource identifier.

### type [ConfigProvider](./config_provider.go#L6)

`type ConfigProvider interface { ... }`

ConfigProvider defines required methods to be used when creating `github_branch_default` terraform resource.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)

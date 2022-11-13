# ghbranchprotect

Package ghbranchprotect provides methods to create `github_branch_protection` terraform resource

## Functions

### func [New](./resource.go#L12)

`func New(c ConfigProvider) *hclwrite.Block`

New returns the `github_branch_protection` terraform resource as `hclwrite.Block`

It returns `nil` if resource is empty.

```golang
valGen := gh2tf.NewValueGenerator()
repoName := "repository_name"
branchPattern := "branch-pattern"
trueBool := "true"
falseBool := "false"
pushRestrictions := []string{"data.github_teams.my_team.node_id"}
contexts := []string{"context"}
dismissalRestrictions := []string{"dismissal-restriction"}
two := "2"
res := &ghbranchprotect.Config{
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
}
hclBlock := ghbranchprotect.New(res)

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(hclBlock)
fmt.Println(string(hclFile.Bytes()))
```

 Output:

```terraform
resource "github_branch_protection" "branch-protection-id" {
  repository_id           = "repository_name"
  pattern                 = "branch-pattern"
  enforce_admins          = true
  allows_deletions        = false
  allows_force_pushes     = true
  push_restrictions       = [data.github_teams.my_team.node_id]
  required_linear_history = false
  require_signed_commits  = true

  required_status_checks {
    strict   = false
    contexts = ["context"]
  }

  required_pull_request_reviews {
    dismiss_stale_reviews           = true
    restrict_dismissals             = false
    dismissal_restrictions          = ["dismissal-restriction"]
    require_code_owner_reviews      = true
    required_approving_review_count = 2
  }
}
```

### func [NewRequiredPRReviewsSignature](./signature.go#L38)

`func NewRequiredPRReviewsSignature(conf RequiredPRReviewsConfigProvider) *tfsig.BlockSignature`

NewRequiredPRReviewsSignature returns the `github_branch_protection->required_pull_request_reviews` terraform block
as `tfsig.BlockSignature`

It returns `nil` if resource is empty.

### func [NewRequiredStatusChecksSignature](./signature.go#L58)

`func NewRequiredStatusChecksSignature(conf RequiredStatusChecksConfigProvider) *tfsig.BlockSignature`

NewRequiredStatusChecksSignature returns the `github_branch_protection->required_status_checks` terraform block
as `tfsig.BlockSignature`

It returns `nil` if resource is empty.

### func [NewSignature](./signature.go#L12)

`func NewSignature(conf ConfigProvider) *tfsig.BlockSignature`

NewSignature returns the `github_branch_protection` terraform resource as `tfsig.BlockSignature`

It returns `nil` if resource is empty.

## Types

### type [Config](./config.go#L12)

`type Config struct { ... }`

Config is the default implementation of `ConfigProvider`.

#### func (*Config) [AllowsDeletionsValue](./config.go#L74)

`func (conf *Config) AllowsDeletionsValue() *cty.Value`

AllowsDeletionsValue return the provided `github_branch_protection` `allow_deletions` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [AllowsForcePushesValue](./config.go#L80)

`func (conf *Config) AllowsForcePushesValue() *cty.Value`

AllowsForcePushesValue return the provided `github_branch_protection` `allows_force_pushes` attribute value
as `cty.Bool` or `nil` if not provided.

#### func (*Config) [EnforceAdminsValue](./config.go#L68)

`func (conf *Config) EnforceAdminsValue() *cty.Value`

EnforceAdminsValue return the provided `github_branch_protection` `enforce_admins` attribute value as `cty.Bool`
or `nil` if not provided.

#### func (*Config) [HasResource](./config.go#L45)

`func (conf *Config) HasResource() bool`

HasResource returns `true` in case at least `Pattern` and `RepositoryId` value exist, else `false`.

#### func (*Config) [PatternValue](./config.go#L62)

`func (conf *Config) PatternValue() *cty.Value`

PatternValue return the provided `github_branch_protection` `pattern` attribute value as `cty.String`
or `nil` if not provided.

#### func (*Config) [PushRestrictionsValue](./config.go#L86)

`func (conf *Config) PushRestrictionsValue() *cty.Value`

PushRestrictionsValue return the provided `github_branch_protection` `push_restrictions` attribute value
as `cty.List` of `cty.String` or `nil` if not provided.

#### func (*Config) [RepositoryIdValue](./config.go#L56)

`func (conf *Config) RepositoryIdValue() *cty.Value`

RepositoryIdValue return the provided `github_branch_protection` `repository_id` attribute value as `cty.String`
or `nil` if not provided.

#### func (*Config) [RequireSignedCommitsValue](./config.go#L98)

`func (conf *Config) RequireSignedCommitsValue() *cty.Value`

RequireSignedCommitsValue return the provided `github_branch_protection` `require_signed_commits` attribute value
as `cty.Bool` or `nil` if not provided.

#### func (*Config) [RequiredLinearHistoryValue](./config.go#L92)

`func (conf *Config) RequiredLinearHistoryValue() *cty.Value`

RequiredLinearHistoryValue return the provided `github_branch_protection` `required_linear_history` attribute value
as `cty.Bool` or `nil` if not provided.

#### func (*Config) [RequiredPullRequestReviewsConfig](./config.go#L108)

`func (conf *Config) RequiredPullRequestReviewsConfig() RequiredPRReviewsConfigProvider`

RequiredPullRequestReviewsConfig return the provided `RequiredPRReviewConfig`.

#### func (*Config) [RequiredStatusChecksConfig](./config.go#L103)

`func (conf *Config) RequiredStatusChecksConfig() RequiredStatusChecksConfigProvider`

RequiredStatusChecksConfig return the provided `RequiredStatusChecksConfig`.

#### func (*Config) [ResourceIdentifier](./config.go#L50)

`func (conf *Config) ResourceIdentifier() string`

ResourceIdentifier returns the provided terraform resource identifier.

### type [ConfigProvider](./config_provider.go#L8)

`type ConfigProvider interface { ... }`

ConfigProvider defines required methods to be used when creating `github_branch_protection` terraform resource.

### type [RequiredPRReviewConfig](./config.go#L35)

`type RequiredPRReviewConfig struct { ... }`

RequiredPRReviewConfig is the default implementation of `RequiredPRReviewConfigProvider`.

#### func (*RequiredPRReviewConfig) [DismissStaleReviewsValue](./config.go#L142)

`func (c *RequiredPRReviewConfig) DismissStaleReviewsValue() *cty.Value`

DismissStaleReviewsValue return the provided `github_branch_protection->required_pull_request_reviews`
`dismiss_stale_reviews` attribute value as `cty.Bool` or `nil` if not provided.

#### func (*RequiredPRReviewConfig) [DismissalRestrictionsValue](./config.go#L154)

`func (c *RequiredPRReviewConfig) DismissalRestrictionsValue() *cty.Value`

DismissalRestrictionsValue return the provided `github_branch_protection->required_pull_request_reviews`
`dismissal_restrictions` attribute value as `cty.List` of `cty.String` or `nil` if not provided.

#### func (*RequiredPRReviewConfig) [HasResource](./config.go#L135)

`func (c *RequiredPRReviewConfig) HasResource() bool`

HasResource returns `true` in case at least one value exist, else `false`.

#### func (*RequiredPRReviewConfig) [RequireCodeOwnerReviewsValue](./config.go#L160)

`func (c *RequiredPRReviewConfig) RequireCodeOwnerReviewsValue() *cty.Value`

RequireCodeOwnerReviewsValue return the provided `github_branch_protection->required_pull_request_reviews`
`require_code_owner_reviews` attribute value as `cty.Bool` or `nil` if not provided.

#### func (*RequiredPRReviewConfig) [RequiredApprovingReviewCountValue](./config.go#L166)

`func (c *RequiredPRReviewConfig) RequiredApprovingReviewCountValue() *cty.Value`

RequiredApprovingReviewCountValue return the provided `github_branch_protection->required_pull_request_reviews`
`required_approving_review_count` attribute value as `cty.Number` or `nil` if not provided.

#### func (*RequiredPRReviewConfig) [RestrictDismissalsValue](./config.go#L148)

`func (c *RequiredPRReviewConfig) RestrictDismissalsValue() *cty.Value`

RestrictDismissalsValue return the provided `github_branch_protection->required_pull_request_reviews`
`restrict_dismissals` attribute value as `cty.Bool` or `nil` if not provided.

### type [RequiredPRReviewsConfigProvider](./config_provider.go#L75)

`type RequiredPRReviewsConfigProvider interface { ... }`

RequiredPRReviewsConfigProvider defines required methods to be used when creating
`github_branch_protection->required_pull_request_reviews` terraform block.

### type [RequiredStatusChecksConfig](./config.go#L28)

`type RequiredStatusChecksConfig struct { ... }`

RequiredStatusChecksConfig is the default implementation of `RequiredStatusChecksConfigProvider`.

#### func (*RequiredStatusChecksConfig) [ContextValue](./config.go#L128)

`func (c *RequiredStatusChecksConfig) ContextValue() *cty.Value`

ContextValue return the provided `github_branch_protection->required_status_checks` `contexts` attribute value
as `cty.List` of `cty.String` or `nil` if not provided.

#### func (*RequiredStatusChecksConfig) [HasResource](./config.go#L116)

`func (c *RequiredStatusChecksConfig) HasResource() bool`

HasResource returns `true` in case at either `RequiredStatusChecksStrict` or `RequiredStatusChecksContext` value
exist, else `false`.

#### func (*RequiredStatusChecksConfig) [StrictValue](./config.go#L122)

`func (c *RequiredStatusChecksConfig) StrictValue() *cty.Value`

StrictValue return the provided `github_branch_protection->required_status_checks` `strict` attribute value
as `cty.Bool` or `nil` if not provided.

### type [RequiredStatusChecksConfigProvider](./config_provider.go#L58)

`type RequiredStatusChecksConfigProvider interface { ... }`

RequiredStatusChecksConfigProvider defines required methods to be used when creating
`github_branch_protection->required_status_checks` terraform block.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)

/*
Package ghbranchprotect provides methods to create `github_branch_protection` terraform resource
*/
package ghbranchprotect

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig"
)

// Config is the default implementation of `ConfigProvider`.
type Config struct {
	ValueGenerator        tfsig.ValueGenerator
	Identifier            string
	RepositoryId          *string
	Pattern               *string
	EnforceAdmins         *string
	AllowsDeletions       *string
	AllowsForcePushes     *string
	PushRestrictions      *[]string
	RequiredLinearHistory *string
	RequireSignedCommits  *string
	RequiredStatusChecks  *RequiredStatusChecksConfig
	RequiredPRReview      *RequiredPRReviewConfig
}

// RequiredStatusChecksConfig is the default implementation of `RequiredStatusChecksConfigProvider`.
type RequiredStatusChecksConfig struct {
	ValueGenerator tfsig.ValueGenerator
	Strict         *string
	Contexts       *[]string
}

// RequiredPRReviewConfig is the default implementation of `RequiredPRReviewConfigProvider`.
type RequiredPRReviewConfig struct {
	ValueGenerator               tfsig.ValueGenerator
	DismissStaleReviews          *string
	RestrictDismissals           *string
	DismissalRestrictions        *[]string
	RequireCodeOwnerReviews      *string
	RequiredApprovingReviewCount *string
}

// HasResource returns `true` in case at least `Pattern` and `RepositoryId` value exist, else `false`.
func (conf *Config) HasResource() bool {
	return conf != nil && conf.Pattern != nil && conf.RepositoryId != nil
}

// ResourceIdentifier returns the provided terraform resource identifier.
func (conf *Config) ResourceIdentifier() string {
	return conf.Identifier
}

// RepositoryIdValue return the provided `github_branch_protection` `repository_id` attribute value as `cty.String`
// or `nil` if not provided.
func (conf *Config) RepositoryIdValue() *cty.Value {
	return conf.ValueGenerator.ToString(conf.RepositoryId)
}

// PatternValue return the provided `github_branch_protection` `pattern` attribute value as `cty.String`
// or `nil` if not provided.
func (conf *Config) PatternValue() *cty.Value {
	return conf.ValueGenerator.ToString(conf.Pattern)
}

// EnforceAdminsValue return the provided `github_branch_protection` `enforce_admins` attribute value as `cty.Bool`
// or `nil` if not provided.
func (conf *Config) EnforceAdminsValue() *cty.Value {
	return conf.ValueGenerator.ToBool(conf.EnforceAdmins)
}

// AllowsDeletionsValue return the provided `github_branch_protection` `allow_deletions` attribute value as `cty.Bool`
// or `nil` if not provided.
func (conf *Config) AllowsDeletionsValue() *cty.Value {
	return conf.ValueGenerator.ToBool(conf.AllowsDeletions)
}

// AllowsForcePushesValue return the provided `github_branch_protection` `allows_force_pushes` attribute value
// as `cty.Bool` or `nil` if not provided.
func (conf *Config) AllowsForcePushesValue() *cty.Value {
	return conf.ValueGenerator.ToBool(conf.AllowsForcePushes)
}

// PushRestrictionsValue return the provided `github_branch_protection` `push_restrictions` attribute value
// as `cty.List` of `cty.String` or `nil` if not provided.
func (conf *Config) PushRestrictionsValue() *cty.Value {
	return conf.ValueGenerator.ToStringList(conf.PushRestrictions)
}

// RequiredLinearHistoryValue return the provided `github_branch_protection` `required_linear_history` attribute value
// as `cty.Bool` or `nil` if not provided.
func (conf *Config) RequiredLinearHistoryValue() *cty.Value {
	return conf.ValueGenerator.ToBool(conf.RequiredLinearHistory)
}

// RequireSignedCommitsValue return the provided `github_branch_protection` `require_signed_commits` attribute value
// as `cty.Bool` or `nil` if not provided.
func (conf *Config) RequireSignedCommitsValue() *cty.Value {
	return conf.ValueGenerator.ToBool(conf.RequireSignedCommits)
}

// RequiredStatusChecksConfig return the provided `RequiredStatusChecksConfig`.
func (conf *Config) RequiredStatusChecksConfig() RequiredStatusChecksConfigProvider { //nolint:ireturn,lll // Disabled as value is passed to a method expecting that interface
	return conf.RequiredStatusChecks
}

// RequiredPullRequestReviewsConfig return the provided `RequiredPRReviewConfig`.
func (conf *Config) RequiredPullRequestReviewsConfig() RequiredPRReviewsConfigProvider { //nolint:ireturn,lll // Disabled as value is passed to a method expecting that interface
	return conf.RequiredPRReview
}

// RequiredStatusChecksConfigProvider implementation

// HasResource returns `true` in case at either `RequiredStatusChecksStrict` or `RequiredStatusChecksContext` value
// exist, else `false`.
func (c *RequiredStatusChecksConfig) HasResource() bool {
	return c != nil && (c.Strict != nil || c.Contexts != nil)
}

// StrictValue return the provided `github_branch_protection->required_status_checks` `strict` attribute value
// as `cty.Bool` or `nil` if not provided.
func (c *RequiredStatusChecksConfig) StrictValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.Strict)
}

// ContextValue return the provided `github_branch_protection->required_status_checks` `contexts` attribute value
// as `cty.List` of `cty.String` or `nil` if not provided.
func (c *RequiredStatusChecksConfig) ContextValue() *cty.Value {
	return c.ValueGenerator.ToStringList(c.Contexts)
}

// RequiredPRReviewsConfigProvider implementation

// HasResource returns `true` in case at least one value exist, else `false`.
func (c *RequiredPRReviewConfig) HasResource() bool {
	return c != nil && (c.DismissStaleReviews != nil || c.RestrictDismissals != nil || c.DismissalRestrictions != nil ||
		c.RequireCodeOwnerReviews != nil || c.RequiredApprovingReviewCount != nil)
}

// DismissStaleReviewsValue return the provided `github_branch_protection->required_pull_request_reviews`
// `dismiss_stale_reviews` attribute value as `cty.Bool` or `nil` if not provided.
func (c *RequiredPRReviewConfig) DismissStaleReviewsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.DismissStaleReviews)
}

// RestrictDismissalsValue return the provided `github_branch_protection->required_pull_request_reviews`
// `restrict_dismissals` attribute value as `cty.Bool` or `nil` if not provided.
func (c *RequiredPRReviewConfig) RestrictDismissalsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.RestrictDismissals)
}

// DismissalRestrictionsValue return the provided `github_branch_protection->required_pull_request_reviews`
// `dismissal_restrictions` attribute value as `cty.List` of `cty.String` or `nil` if not provided.
func (c *RequiredPRReviewConfig) DismissalRestrictionsValue() *cty.Value {
	return c.ValueGenerator.ToStringList(c.DismissalRestrictions)
}

// RequireCodeOwnerReviewsValue return the provided `github_branch_protection->required_pull_request_reviews`
// `require_code_owner_reviews` attribute value as `cty.Bool` or `nil` if not provided.
func (c *RequiredPRReviewConfig) RequireCodeOwnerReviewsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.RequireCodeOwnerReviews)
}

// RequiredApprovingReviewCountValue return the provided `github_branch_protection->required_pull_request_reviews`
// `required_approving_review_count` attribute value as `cty.Number` or `nil` if not provided.
func (c *RequiredPRReviewConfig) RequiredApprovingReviewCountValue() *cty.Value {
	return c.ValueGenerator.ToNumber(c.RequiredApprovingReviewCount)
}

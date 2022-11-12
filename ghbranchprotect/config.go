/*
Package ghbranchprotect provides methods to create `github_branch_protection` terraform resource
*/
package ghbranchprotect

import (
	"github.com/yoanm/go-tfsig"
	"github.com/zclconf/go-cty/cty"
)

// Config is the default implementation of `ConfigProvider`
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

// RequiredStatusChecksConfig is the default implementation of `RequiredStatusChecksConfigProvider`
type RequiredStatusChecksConfig struct {
	ValueGenerator tfsig.ValueGenerator
	Strict         *string
	Contexts       *[]string
}

// RequiredPRReviewConfig is the default implementation of `RequiredPRReviewConfigProvider`
type RequiredPRReviewConfig struct {
	ValueGenerator               tfsig.ValueGenerator
	DismissStaleReviews          *string
	RestrictDismissals           *string
	DismissalRestrictions        *[]string
	RequireCodeOwnerReviews      *string
	RequiredApprovingReviewCount *string
}

// HasResource returns `true` in case at least `Pattern` and `RepositoryId` value exist, else `false`
func (c *Config) HasResource() bool {
	return c != nil && c.Pattern != nil && c.RepositoryId != nil
}

// ResourceIdentifier returns the provided terraform resource identifier
func (c *Config) ResourceIdentifier() string {
	return c.Identifier
}

// RepositoryIdValue return the provided `github_branch_protection` `repository_id` attribute value as `cty.String` or `nil` if not provided
func (c *Config) RepositoryIdValue() *cty.Value {
	return c.ValueGenerator.ToString(c.RepositoryId)
}

// PatternValue return the provided `github_branch_protection` `pattern` attribute value as `cty.String` or `nil` if not provided
func (c *Config) PatternValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Pattern)
}

// EnforceAdminsValue return the provided `github_branch_protection` `enforce_admins` attribute value as `cty.Bool` or `nil` if not provided
func (c *Config) EnforceAdminsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.EnforceAdmins)
}

// AllowsDeletionsValue return the provided `github_branch_protection` `allow_deletions` attribute value as `cty.Bool` or `nil` if not provided
func (c *Config) AllowsDeletionsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.AllowsDeletions)
}

// AllowsForcePushesValue return the provided `github_branch_protection` `allows_force_pushes` attribute value as `cty.Bool` or `nil` if not provided
func (c *Config) AllowsForcePushesValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.AllowsForcePushes)
}

// PushRestrictionsValue return the provided `github_branch_protection` `push_restrictions` attribute value as `cty.List` of `cty.String` or `nil` if not provided
func (c *Config) PushRestrictionsValue() *cty.Value {
	return c.ValueGenerator.ToStringList(c.PushRestrictions)
}

// RequiredLinearHistoryValue return the provided `github_branch_protection` `required_linear_history` attribute value as `cty.Bool` or `nil` if not provided
func (c *Config) RequiredLinearHistoryValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.RequiredLinearHistory)
}

// RequireSignedCommitsValue return the provided `github_branch_protection` `require_signed_commits` attribute value as `cty.Bool` or `nil` if not provided
func (c *Config) RequireSignedCommitsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.RequireSignedCommits)
}

// RequiredStatusChecksConfig return the provided `RequiredStatusChecksConfig`
func (c *Config) RequiredStatusChecksConfig() RequiredStatusChecksConfigProvider {
	return c.RequiredStatusChecks
}

// RequiredPullRequestReviewsConfig return the provided `RequiredPRReviewConfig`
func (c *Config) RequiredPullRequestReviewsConfig() RequiredPRReviewsConfigProvider {
	return c.RequiredPRReview
}

// RequiredStatusChecksConfigProvider implementation

// HasResource returns `true` in case at either `RequiredStatusChecksStrict` or `RequiredStatusChecksContext` value exist, else `false`
func (c *RequiredStatusChecksConfig) HasResource() bool {
	return c != nil && (c.Strict != nil || c.Contexts != nil)
}

// StrictValue return the provided `github_branch_protection->required_status_checks` `strict` attribute value as `cty.Bool` or `nil` if not provided
func (c *RequiredStatusChecksConfig) StrictValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.Strict)
}

// ContextValue return the provided `github_branch_protection->required_status_checks` `contexts` attribute value as `cty.List` of `cty.String` or `nil` if not provided
func (c *RequiredStatusChecksConfig) ContextValue() *cty.Value {
	return c.ValueGenerator.ToStringList(c.Contexts)
}

// RequiredPRReviewsConfigProvider implementation

// HasResource returns `true` in case at least one value exist, else `false`
func (c *RequiredPRReviewConfig) HasResource() bool {
	return c != nil && (c.DismissStaleReviews != nil || c.RestrictDismissals != nil || c.DismissalRestrictions != nil ||
		c.RequireCodeOwnerReviews != nil || c.RequiredApprovingReviewCount != nil)
}

// DismissStaleReviewsValue return the provided `github_branch_protection->required_pull_request_reviews` `dismiss_stale_reviews` attribute value as `cty.Bool` or `nil` if not provided
func (c *RequiredPRReviewConfig) DismissStaleReviewsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.DismissStaleReviews)
}

// RestrictDismissalsValue return the provided `github_branch_protection->required_pull_request_reviews` `restrict_dismissals` attribute value as `cty.Bool` or `nil` if not provided
func (c *RequiredPRReviewConfig) RestrictDismissalsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.RestrictDismissals)
}

// DismissalRestrictionsValue return the provided `github_branch_protection->required_pull_request_reviews` `dismissal_restrictions` attribute value as `cty.List` of `cty.String` or `nil` if not provided
func (c *RequiredPRReviewConfig) DismissalRestrictionsValue() *cty.Value {
	return c.ValueGenerator.ToStringList(c.DismissalRestrictions)
}

// RequireCodeOwnerReviewsValue return the provided `github_branch_protection->required_pull_request_reviews` `require_code_owner_reviews` attribute value as `cty.Bool` or `nil` if not provided
func (c *RequiredPRReviewConfig) RequireCodeOwnerReviewsValue() *cty.Value {
	return c.ValueGenerator.ToBool(c.RequireCodeOwnerReviews)
}

// RequiredApprovingReviewCountValue return the provided `github_branch_protection->required_pull_request_reviews` `required_approving_review_count` attribute value as `cty.Number` or `nil` if not provided
func (c *RequiredPRReviewConfig) RequiredApprovingReviewCountValue() *cty.Value {
	return c.ValueGenerator.ToNumber(c.RequiredApprovingReviewCount)
}

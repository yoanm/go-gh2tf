package ghbranchprotect

import "github.com/zclconf/go-cty/cty"

// ConfigProvider defines required methods to be used when creating `github_branch_protection` terraform resource
type ConfigProvider interface {
	// HasResource will indicate if resource has to be created (return `true` for instance to always create the resource even if empty)
	HasResource() bool
	// ResourceIdentifier returns the resource identifier
	ResourceIdentifier() string

	// RepositoryIdValue return the `github_branch_protection` `repository_id` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	RepositoryIdValue() *cty.Value
	// PatternValue return the `github_branch_protection` `pattern` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	PatternValue() *cty.Value
	// EnforceAdminsValue return the `github_branch_protection` `enforce_admins` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	EnforceAdminsValue() *cty.Value
	// AllowsDeletionsValue return the `github_branch_protection` `allow_deletions` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	AllowsDeletionsValue() *cty.Value
	// AllowsForcePushesValue return the `github_branch_protection` `allows_force_pushes` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	AllowsForcePushesValue() *cty.Value
	// PushRestrictionsValue return the `github_branch_protection` `push_restrictions` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	PushRestrictionsValue() *cty.Value
	// RequiredLinearHistoryValue return the `github_branch_protection` `required_linear_history` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	RequiredLinearHistoryValue() *cty.Value
	// RequireSignedCommitsValue return the `github_branch_protection` `require_signed_commits` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	RequireSignedCommitsValue() *cty.Value
	// RequiredStatusChecksConfig return the config provider for `github_branch_protection->required_status_checks` block
	RequiredStatusChecksConfig() RequiredStatusChecksConfigProvider
	// RequiredPullRequestReviewsConfig return the config provider for `github_branch_protection->required_pull_request_reviews` block
	RequiredPullRequestReviewsConfig() RequiredPRReviewsConfigProvider
}

// RequiredStatusChecksConfigProvider defines required methods to be used when creating `github_branch_protection->required_status_checks` terraform block
type RequiredStatusChecksConfigProvider interface {
	// HasResource will indicate if resource has to be created (return `true` for instance to always create the resource even if empty)
	HasResource() bool

	// StrictValue return the `github_branch_protection->required_status_checks` `strict` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	StrictValue() *cty.Value
	// ContextValue return the `github_branch_protection->required_status_checks` `contexts` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	ContextValue() *cty.Value
}

// RequiredPRReviewsConfigProvider defines required methods to be used when creating `github_branch_protection->required_pull_request_reviews` terraform block
type RequiredPRReviewsConfigProvider interface {
	// HasResource will indicate if resource has to be created (return `true` for instance to always create the resource even if empty)
	HasResource() bool

	// DismissStaleReviewsValue return the `github_branch_protection->required_pull_request_reviews` `dismiss_stale_reviews` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	DismissStaleReviewsValue() *cty.Value
	// RestrictDismissalsValue return the `github_branch_protection->required_pull_request_reviews` `restrict_dismissals` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	RestrictDismissalsValue() *cty.Value
	// DismissalRestrictionsValue return the `github_branch_protection->required_pull_request_reviews` `dismissal_restrictions` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	DismissalRestrictionsValue() *cty.Value
	// RequireCodeOwnerReviewsValue return the `github_branch_protection->required_pull_request_reviews` `require_code_owner_reviews` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	RequireCodeOwnerReviewsValue() *cty.Value
	// RequiredApprovingReviewCountValue return the `github_branch_protection->required_pull_request_reviews` `required_approving_review_count` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	RequiredApprovingReviewCountValue() *cty.Value
}

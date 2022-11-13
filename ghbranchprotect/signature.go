package ghbranchprotect

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

	sig := tfsig.NewEmptyResource("github_branch_protection", conf.ResourceIdentifier())

	appendAttrIfNotNil(sig, "repository_id", conf.RepositoryIdValue())
	appendAttrIfNotNil(sig, "pattern", conf.PatternValue())
	appendAttrIfNotNil(sig, "enforce_admins", conf.EnforceAdminsValue())
	appendAttrIfNotNil(sig, "allows_deletions", conf.AllowsDeletionsValue())
	appendAttrIfNotNil(sig, "allows_force_pushes", conf.AllowsForcePushesValue())
	appendAttrIfNotNil(sig, "push_restrictions", conf.PushRestrictionsValue())
	appendAttrIfNotNil(sig, "required_linear_history", conf.RequiredLinearHistoryValue())
	appendAttrIfNotNil(sig, "require_signed_commits", conf.RequireSignedCommitsValue())

	appendChildIfNotNil(sig, NewRequiredStatusChecksSignature(conf.RequiredStatusChecksConfig()))
	appendChildIfNotNil(sig, NewRequiredPRReviewsSignature(conf.RequiredPullRequestReviewsConfig()))

	return sig
}

// NewRequiredPRReviewsSignature returns the `github_branch_protection->required_pull_request_reviews` terraform block
// as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewRequiredPRReviewsSignature(conf RequiredPRReviewsConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptySignature("required_pull_request_reviews")

	appendAttrIfNotNil(sig, "dismiss_stale_reviews", conf.DismissStaleReviewsValue())
	appendAttrIfNotNil(sig, "restrict_dismissals", conf.RestrictDismissalsValue())
	appendAttrIfNotNil(sig, "dismissal_restrictions", conf.DismissalRestrictionsValue())
	appendAttrIfNotNil(sig, "require_code_owner_reviews", conf.RequireCodeOwnerReviewsValue())
	appendAttrIfNotNil(sig, "required_approving_review_count", conf.RequiredApprovingReviewCountValue())

	return sig
}

// NewRequiredStatusChecksSignature returns the `github_branch_protection->required_status_checks` terraform block
// as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewRequiredStatusChecksSignature(conf RequiredStatusChecksConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptySignature("required_status_checks")

	appendAttrIfNotNil(sig, "strict", conf.StrictValue())
	appendAttrIfNotNil(sig, "contexts", conf.ContextValue())

	return sig
}

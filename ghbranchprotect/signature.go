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

	sig := tfsig.NewResource("github_branch_protection", conf.ResourceIdentifier())

	tfsig.AppendAttributeIfNotNil(sig, "repository_id", conf.RepositoryIdValue())
	tfsig.AppendAttributeIfNotNil(sig, "pattern", conf.PatternValue())
	tfsig.AppendAttributeIfNotNil(sig, "enforce_admins", conf.EnforceAdminsValue())
	tfsig.AppendAttributeIfNotNil(sig, "allows_deletions", conf.AllowsDeletionsValue())
	tfsig.AppendAttributeIfNotNil(sig, "allows_force_pushes", conf.AllowsForcePushesValue())
	tfsig.AppendAttributeIfNotNil(sig, "push_restrictions", conf.PushRestrictionsValue())
	tfsig.AppendAttributeIfNotNil(sig, "required_linear_history", conf.RequiredLinearHistoryValue())
	tfsig.AppendAttributeIfNotNil(sig, "require_signed_commits", conf.RequireSignedCommitsValue())

	tfsig.AppendChildIfNotNil(sig, NewRequiredStatusChecksSignature(conf.RequiredStatusChecksConfig()))
	tfsig.AppendChildIfNotNil(sig, NewRequiredPRReviewsSignature(conf.RequiredPullRequestReviewsConfig()))

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

	sig := tfsig.NewSignature("required_pull_request_reviews")

	tfsig.AppendAttributeIfNotNil(sig, "dismiss_stale_reviews", conf.DismissStaleReviewsValue())
	tfsig.AppendAttributeIfNotNil(sig, "restrict_dismissals", conf.RestrictDismissalsValue())
	tfsig.AppendAttributeIfNotNil(sig, "dismissal_restrictions", conf.DismissalRestrictionsValue())
	tfsig.AppendAttributeIfNotNil(sig, "require_code_owner_reviews", conf.RequireCodeOwnerReviewsValue())
	tfsig.AppendAttributeIfNotNil(sig, "required_approving_review_count", conf.RequiredApprovingReviewCountValue())

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

	sig := tfsig.NewSignature("required_status_checks")

	tfsig.AppendAttributeIfNotNil(sig, "strict", conf.StrictValue())
	tfsig.AppendAttributeIfNotNil(sig, "contexts", conf.ContextsValue())

	return sig
}

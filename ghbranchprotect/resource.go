package ghbranchprotect

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-tfsig"
)

// tell PMD-CPD to start ignoring code as it is duplicated over gh* packages (look for "CPD-ON" comment below) - CPD-OFF

// New returns the `github_branch_protection` terraform resource as `hclwrite.Block`
//
// It returns `nil` if resource is empty.
func New(c ConfigProvider) *hclwrite.Block {
	if s := NewSignature(c); s != nil {
		return s.Build()
	}

	return nil
}

// NewSignature returns the `github_branch_protection` terraform resource as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
//
//nolint:cyclop // disabled as complexity is based on number of terraform attributes
func NewSignature(conf ConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptyResource("github_branch_protection", conf.ResourceIdentifier())
	// resume PMD-CPD analysis - CPD-ON

	if v := conf.RepositoryIdValue(); v != nil {
		sig.AppendAttribute("repository_id", *v)
	}

	if v := conf.PatternValue(); v != nil {
		sig.AppendAttribute("pattern", *v)
	}

	if v := conf.EnforceAdminsValue(); v != nil {
		sig.AppendAttribute("enforce_admins", *v)
	}

	if v := conf.AllowsDeletionsValue(); v != nil {
		sig.AppendAttribute("allows_deletions", *v)
	}

	if v := conf.AllowsForcePushesValue(); v != nil {
		sig.AppendAttribute("allows_force_pushes", *v)
	}

	if v := conf.PushRestrictionsValue(); v != nil {
		sig.AppendAttribute("push_restrictions", *v)
	}

	if v := conf.RequiredLinearHistoryValue(); v != nil {
		sig.AppendAttribute("required_linear_history", *v)
	}

	if v := conf.RequireSignedCommitsValue(); v != nil {
		sig.AppendAttribute("require_signed_commits", *v)
	}

	if s := NewRequiredStatusChecksSignature(conf.RequiredStatusChecksConfig()); s != nil {
		sig.AppendEmptyLine()
		sig.AppendChild(s)
	}

	if s := NewRequiredPRReviewsSignature(conf.RequiredPullRequestReviewsConfig()); s != nil {
		sig.AppendEmptyLine()
		sig.AppendChild(s)
	}

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

	if v := conf.DismissStaleReviewsValue(); v != nil {
		sig.AppendAttribute("dismiss_stale_reviews", *v)
	}

	if v := conf.RestrictDismissalsValue(); v != nil {
		sig.AppendAttribute("restrict_dismissals", *v)
	}

	if v := conf.DismissalRestrictionsValue(); v != nil {
		sig.AppendAttribute("dismissal_restrictions", *v)
	}

	if v := conf.RequireCodeOwnerReviewsValue(); v != nil {
		sig.AppendAttribute("require_code_owner_reviews", *v)
	}

	if v := conf.RequiredApprovingReviewCountValue(); v != nil {
		sig.AppendAttribute("required_approving_review_count", *v)
	}

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

	if v := conf.StrictValue(); v != nil {
		sig.AppendAttribute("strict", *v)
	}

	if v := conf.ContextValue(); v != nil {
		sig.AppendAttribute("contexts", *v)
	}

	return sig
}

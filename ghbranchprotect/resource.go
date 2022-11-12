package ghbranchprotect

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-tfsig"
)

// New returns the `github_branch_protection` terraform resource as `hclwrite.Block`
//
// It returns `nil` if resource is empty
func New(c ConfigProvider) *hclwrite.Block {
	if s := NewSignature(c); s != nil {
		return s.Build()
	}

	return nil
}

// NewSignature returns the `github_branch_protection` terraform resource as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty
func NewSignature(c ConfigProvider) *tfsig.BlockSignature {
	if c == nil || !c.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptyResource("github_branch_protection", c.ResourceIdentifier())

	if v := c.RepositoryIdValue(); v != nil {
		sig.AppendAttribute("repository_id", *v)
	}
	if v := c.PatternValue(); v != nil {
		sig.AppendAttribute("pattern", *v)
	}
	if v := c.EnforceAdminsValue(); v != nil {
		sig.AppendAttribute("enforce_admins", *v)
	}
	if v := c.AllowsDeletionsValue(); v != nil {
		sig.AppendAttribute("allows_deletions", *v)
	}
	if v := c.AllowsForcePushesValue(); v != nil {
		sig.AppendAttribute("allows_force_pushes", *v)
	}
	if v := c.PushRestrictionsValue(); v != nil {
		sig.AppendAttribute("push_restrictions", *v)
	}
	if v := c.RequiredLinearHistoryValue(); v != nil {
		sig.AppendAttribute("required_linear_history", *v)
	}
	if v := c.RequireSignedCommitsValue(); v != nil {
		sig.AppendAttribute("require_signed_commits", *v)
	}

	if v := c.RequiredStatusChecksConfig(); v != nil {
		if s := NewRequiredStatusChecksSignature(v); s != nil {
			sig.AppendEmptyLine()
			sig.AppendChild(s)
		}
	}

	if v := c.RequiredPullRequestReviewsConfig(); v != nil {
		if s := NewRequiredPRReviewsSignature(v); s != nil {
			sig.AppendEmptyLine()
			sig.AppendChild(s)
		}
	}

	return sig
}

// NewRequiredPRReviewsSignature returns the `github_branch_protection->required_pull_request_reviews` terraform block as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty
func NewRequiredPRReviewsSignature(c RequiredPRReviewsConfigProvider) *tfsig.BlockSignature {
	if c == nil || !c.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptySignature("required_pull_request_reviews")

	if v := c.DismissStaleReviewsValue(); v != nil {
		sig.AppendAttribute("dismiss_stale_reviews", *v)
	}
	if v := c.RestrictDismissalsValue(); v != nil {
		sig.AppendAttribute("restrict_dismissals", *v)
	}
	if v := c.DismissalRestrictionsValue(); v != nil {
		sig.AppendAttribute("dismissal_restrictions", *v)
	}
	if v := c.RequireCodeOwnerReviewsValue(); v != nil {
		sig.AppendAttribute("require_code_owner_reviews", *v)
	}
	if v := c.RequiredApprovingReviewCountValue(); v != nil {
		sig.AppendAttribute("required_approving_review_count", *v)
	}

	return sig
}

// NewRequiredStatusChecksSignature returns the `github_branch_protection->required_status_checks` terraform block as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty
func NewRequiredStatusChecksSignature(c RequiredStatusChecksConfigProvider) *tfsig.BlockSignature {
	if c == nil || !c.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptySignature("required_status_checks")

	if v := c.StrictValue(); v != nil {
		sig.AppendAttribute("strict", *v)
	}
	if v := c.ContextValue(); v != nil {
		sig.AppendAttribute("contexts", *v)
	}

	return sig
}

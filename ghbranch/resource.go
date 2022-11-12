package ghbranch

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-tfsig"
)

// New returns the `github_branch` terraform resource as `hclwrite.Block`
//
// It returns `nil` if `NewSignature` returns nil
func New(c ConfigProvider) *hclwrite.Block {
	if signature := NewSignature(c); signature != nil {
		return signature.Build()
	}

	return nil
}

// NewSignature returns the `github_branch` terraform resource as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty
func NewSignature(c ConfigProvider) *tfsig.BlockSignature {
	if c == nil || !c.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptyResource("github_branch", c.ResourceIdentifier())

	if v := c.RepositoryValue(); v != nil {
		sig.AppendAttribute("repository", *v)
	}
	if v := c.BranchValue(); v != nil {
		sig.AppendAttribute("branch", *v)
	}
	if v := c.SourceBranchValue(); v != nil {
		sig.AppendAttribute("source_branch", *v)
	}

	if v := c.SourceShaValue(); v != nil {
		sig.AppendAttribute("source_sha", *v)
	}

	return sig
}

package ghbranchdefault

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-tfsig"
)

// New returns the `github_branch_default` terraform resource as `hclwrite.Block`
//
// It returns `nil` if resource is empty
func New(c ConfigProvider) *hclwrite.Block {
	if signature := NewSignature(c); signature != nil {
		return signature.Build()
	}

	return nil
}

// NewSignature returns the `github_branch_default` terraform resource as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty
func NewSignature(c ConfigProvider) *tfsig.BlockSignature {
	if c == nil || !c.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptyResource("github_branch_default", c.ResourceIdentifier())

	if v := c.RepositoryValue(); v != nil {
		sig.AppendAttribute("repository", *v)
	}
	if v := c.BranchValue(); v != nil {
		sig.AppendAttribute("branch", *v)
	}

	return sig
}

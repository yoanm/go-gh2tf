package ghbranchdefault

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-tfsig"
)

// New returns the `github_branch_default` terraform resource as `hclwrite.Block`
//
// It returns `nil` if resource is empty.
func New(c ConfigProvider) *hclwrite.Block {
	if signature := NewSignature(c); signature != nil {
		return signature.Build()
	}

	return nil
}

// NewSignature returns the `github_branch_default` terraform resource as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewSignature(conf ConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptyResource("github_branch_default", conf.ResourceIdentifier())

	if v := conf.RepositoryValue(); v != nil {
		sig.AppendAttribute("repository", *v)
	}

	if v := conf.BranchValue(); v != nil {
		sig.AppendAttribute("branch", *v)
	}

	return sig
}

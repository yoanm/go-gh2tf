package ghbranch

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-tfsig"
)

// tell PMD-CPD to start ignoring code as it is duplicated over gh* packages (look for "CPD-ON" comment below) - CPD-OFF

// New returns the `github_branch` terraform resource as `hclwrite.Block`
//
// It returns `nil` if resource is empty.
func New(c ConfigProvider) *hclwrite.Block {
	if signature := NewSignature(c); signature != nil {
		return signature.Build()
	}

	return nil
}

// NewSignature returns the `github_branch` terraform resource as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewSignature(conf ConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptyResource("github_branch", conf.ResourceIdentifier())
	// resume PMD-CPD analysis - CPD-ON

	if v := conf.RepositoryValue(); v != nil {
		sig.AppendAttribute("repository", *v)
	}

	if v := conf.BranchValue(); v != nil {
		sig.AppendAttribute("branch", *v)
	}

	if v := conf.SourceBranchValue(); v != nil {
		sig.AppendAttribute("source_branch", *v)
	}

	if v := conf.SourceShaValue(); v != nil {
		sig.AppendAttribute("source_sha", *v)
	}

	return sig
}

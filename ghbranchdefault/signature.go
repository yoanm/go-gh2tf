package ghbranchdefault

import (
	"github.com/yoanm/go-tfsig"
)

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

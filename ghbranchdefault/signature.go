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

	sig := tfsig.NewResource("github_branch_default", conf.ResourceIdentifier())

	tfsig.AppendAttributeIfNotNil(sig, "repository", conf.RepositoryValue())
	tfsig.AppendAttributeIfNotNil(sig, "branch", conf.BranchValue())

	return sig
}

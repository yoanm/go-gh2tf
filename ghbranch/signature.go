package ghbranch

import (
	"github.com/yoanm/go-tfsig"
)

// NewSignature returns the `github_branch` terraform resource as `tfsig.BlockSignature`
//
// It returns `nil` if resource is empty.
func NewSignature(conf ConfigProvider) *tfsig.BlockSignature {
	if conf == nil || !conf.HasResource() {
		return nil
	}

	sig := tfsig.NewEmptyResource("github_branch", conf.ResourceIdentifier())

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

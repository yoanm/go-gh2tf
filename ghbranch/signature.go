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

	sig := tfsig.NewResource("github_branch", conf.ResourceIdentifier())

	tfsig.AppendAttributeIfNotNil(sig, "repository", conf.RepositoryValue())
	tfsig.AppendAttributeIfNotNil(sig, "branch", conf.BranchValue())
	tfsig.AppendAttributeIfNotNil(sig, "source_branch", conf.SourceBranchValue())
	tfsig.AppendAttributeIfNotNil(sig, "source_sha", conf.SourceShaValue())

	return sig
}

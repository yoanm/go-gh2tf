package ghbranch

import "github.com/zclconf/go-cty/cty"

// ConfigProvider defines required methods to be used when creating `github_branch` terraform resource
type ConfigProvider interface {
	// HasResource will indicate if resource has to be created (return `true` for instance to always create the resource even if empty)
	HasResource() bool
	// ResourceIdentifier returns the resource identifier
	ResourceIdentifier() string

	// RepositoryValue return the `github_branch` `repository` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	RepositoryValue() *cty.Value
	// BranchValue return the `github_branch` `branch` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	BranchValue() *cty.Value
	// SourceBranchValue return the `github_branch` `source_branch` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	SourceBranchValue() *cty.Value
	// SourceShaValue return the `github_branch` `source_sha` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	SourceShaValue() *cty.Value
}

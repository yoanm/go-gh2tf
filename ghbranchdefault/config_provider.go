package ghbranchdefault

import "github.com/zclconf/go-cty/cty"

// ConfigProvider defines required methods to be used when creating `github_branch_default` terraform resource.
type ConfigProvider interface {
	// HasResource will indicate if resource has to be created (return `true` for instance to always
	// create the resource even if empty)
	HasResource() bool
	// ResourceIdentifier returns the resource identifier
	ResourceIdentifier() string

	// RepositoryValue return the `github_branch_default` `repository` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	RepositoryValue() *cty.Value
	// BranchValue return the `github_branch_default` `branch` attribute value.
	//
	// If `nil` is returned, attribute will be omitted
	BranchValue() *cty.Value
}

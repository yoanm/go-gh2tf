/*
Package ghbranchdefault provides methods to create `github_branch_default` terraform resource
*/
package ghbranchdefault

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig"
)

// Config is the default implementation of `ConfigProvider`.
type Config struct {
	ValueGenerator tfsig.ValueGenerator
	Identifier     string
	Repository     *string
	Branch         *string
}

// HasResource returns `true` in case both `BranchConfig` and `Repository` value exist, else `false`.
func (c *Config) HasResource() bool {
	return c != nil && c.Branch != nil && c.Repository != nil
}

// ResourceIdentifier returns the provided terraform resource identifier.
func (c *Config) ResourceIdentifier() string {
	return c.Identifier
}

// RepositoryValue return the provided `github_branch_default` `repository` attribute value as `cty.String`
// or `nil` if not provided.
func (c *Config) RepositoryValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Repository)
}

// BranchValue return the provided `github_branch_default` `branch` attribute value as `cty.String`
// or `nil` if not provided.
func (c *Config) BranchValue() *cty.Value {
	return c.ValueGenerator.ToString(c.Branch)
}

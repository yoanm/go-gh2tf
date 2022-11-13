package ghbranchprotect

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

/** Public **/

// New returns the `github_branch_protection` terraform resource as `hclwrite.Block`
//
// It returns `nil` if resource is empty.
func New(c ConfigProvider) *hclwrite.Block {
	if s := NewSignature(c); s != nil {
		return s.Build()
	}

	return nil
}

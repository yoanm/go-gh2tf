package ghbranchdefault

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
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

package ghrepository

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

/** Public **/

// New returns the `github_repository` terraform resource as `hclwrite.Block`
//
// It returns `nil` if resource is empty.
func New(conf ConfigProvider) *hclwrite.Block {
	if signature := NewSignature(conf); signature != nil {
		return signature.Build()
	}

	return nil
}

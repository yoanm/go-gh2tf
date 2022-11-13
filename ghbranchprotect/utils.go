package ghbranchprotect

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig"
)

func appendAttrIfNotNil(sig *tfsig.BlockSignature, attrName string, v *cty.Value) {
	if v != nil {
		sig.AppendAttribute(attrName, *v)
	}
}

func appendChildIfNotNil(sig *tfsig.BlockSignature, child *tfsig.BlockSignature) {
	if child != nil {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		sig.AppendChild(child)
	}
}

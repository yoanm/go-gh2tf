package ghbranchprotect

import (
	"github.com/yoanm/go-tfsig"
	"github.com/zclconf/go-cty/cty"
)

func appendAttrIfNotNil(sig *tfsig.BlockSignature, attrName string, v *cty.Value) {
	if v != nil {
		sig.AppendAttribute(attrName, *v)
	}
}

func appendBlockAndEmptyLineIfNotNil(sig *tfsig.BlockSignature, b *tfsig.BlockSignature) {
	if b != nil {
		sig.AppendEmptyLine()
		sig.AppendChild(b)
	}
}

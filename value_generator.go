package gh2tf

import "github.com/yoanm/go-tfsig"

// NewValueGenerator returns a new ValueGenerator with provided list of prefix to consider as 'ident' tokens
//
// Default `tfsig` ident tokens and terraform gitHub resources will be automatically added.
func NewValueGenerator(identPrefixList ...string) tfsig.ValueGenerator {
	return tfsig.NewValueGeneratorWith(NewIdentTokenMatcher(identPrefixList...))
}

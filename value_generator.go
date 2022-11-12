package gh2tf

import "github.com/yoanm/go-tfsig"

// NewValueGenerator returns a new ValueGenerator with all terraform gitHub resources escaped to consider as 'ident' tokens
func NewValueGenerator() tfsig.ValueGenerator {
	return tfsig.NewValueGeneratorWith(NewIdentTokenMatcher())
}

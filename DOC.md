# gh2tf

## Functions

### func [NewIdentTokenMatcher](/token_ident_matcher.go#L52)

`func NewIdentTokenMatcher(prefixList ...string) tfsig.IdentTokenMatcher`

NewIdentTokenMatcher returns an instance of IdentTokenMatcher with provided list of prefix to consider as 'ident' tokens

Default `tfsig` ident tokens and terraform gitHub resources will be automatically added

### func [NewValueGenerator](/value_generator.go#L6)

`func NewValueGenerator() tfsig.ValueGenerator`

NewValueGenerator returns a new ValueGenerator with all terraform gitHub resources escaped to consider as 'ident' tokens

## Sub Packages

* [ghbranch](./ghbranch)

* [ghbranchdefault](./ghbranchdefault)

* [ghbranchprotect](./ghbranchprotect)

* [ghrepository](./ghrepository)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)

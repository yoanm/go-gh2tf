# gh2tf

Package gh2tf is an helper for terraform gitHub resources

It provides methods to easily create GitHub HCL resources (`github_repository`, `github_branch`, etc)

## Functions

### func [NewIdentTokenMatcher](/token_ident_matcher.go#L57)

`func NewIdentTokenMatcher(prefixList ...string) tfsig.IdentTokenMatcher`

NewIdentTokenMatcher returns an instance of IdentTokenMatcher with provided list of prefix to consider as 'ident' tokens

Default `tfsig` ident tokens and terraform gitHub resources will be automatically added

### func [NewValueGenerator](/value_generator.go#L6)

`func NewValueGenerator(identPrefixList ...string) tfsig.ValueGenerator`

NewValueGenerator returns a new ValueGenerator with all terraform gitHub resources escaped to consider as 'ident' tokens

```golang
basicStringValue := "basic_value"
localVal := "local.my_local"
varVal := "var.my_var"
dataVal := "data.my_data.my_property"
customStringValue := "custom.my_var"
identStingValue := "github_repository.res-id.name"
identListStringValue := []string{"github_branch.res-id.branch", "github_branch_protection.res-id.pattern"}
explicitIdentStringValue := "explicit_ident.foo"
explicitIdentListStringValue := []string{"explicit_ident_item.foo", "explicit_ident_item.bar"}

valGen := NewValueGenerator()
sig := tfsig.NewEmptySignature("my_block")
sig.AppendAttribute("attr1", *valGen.ToString(&basicStringValue))
sig.AppendAttribute("attr2", *valGen.ToString(&localVal))
sig.AppendAttribute("attr3", *valGen.ToString(&varVal))
sig.AppendAttribute("attr4", *valGen.ToString(&dataVal))
sig.AppendAttribute("attr5", *valGen.ToString(&customStringValue))
sig.AppendAttribute("attr6", *valGen.ToString(&identStingValue))
sig.AppendAttribute("attr7", *valGen.ToStringList(&identListStringValue))
sig.AppendEmptyLine()
sig.AppendAttribute("attr8", *valGen.ToIdent(&explicitIdentStringValue))
sig.AppendAttribute("attr9", *valGen.ToIdentList(&explicitIdentListStringValue))
customValGen := NewValueGenerator("custom.")
sig.AppendEmptyLine()
sig.AppendAttribute("attr10", *customValGen.ToString(&customStringValue))

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(sig.Build())
fmt.Println(string(hclFile.Bytes()))
```

 Output:

```
my_block {
  attr1 = "basic_value"
  attr2 = local.my_local
  attr3 = var.my_var
  attr4 = data.my_data.my_property
  attr5 = "custom.my_var"
  attr6 = github_repository.res-id.name
  attr7 = [github_branch.res-id.branch, github_branch_protection.res-id.pattern]

  attr8 = explicit_ident.foo
  attr9 = [explicit_ident_item.foo, explicit_ident_item.bar]

  attr10 = custom.my_var
}
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)

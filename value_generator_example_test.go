package gh2tf_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/yoanm/go-tfsig"

	"github.com/yoanm/go-gh2tf"
)

func ExampleNewValueGenerator() {
	basicStringValue := "basic_value"
	localVal := "local.my_local"
	varVal := "var.my_var"
	dataVal := "data.my_data.my_property"
	customStringValue := "custom.my_var"
	identStingValue := "github_repository.res-id.name"
	identListStringValue := []string{"github_branch.res-id.branch", "github_branch_protection.res-id.pattern"}
	explicitIdentStringValue := "explicit_ident.foo"
	explicitIdentListStringValue := []string{"explicit_ident_item.foo", "explicit_ident_item.bar"}

	valGen := gh2tf.NewValueGenerator()
	sig := tfsig.NewSignature("my_block")
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

	customValGen := gh2tf.NewValueGenerator("custom.")

	sig.AppendEmptyLine()
	sig.AppendAttribute("attr10", *customValGen.ToString(&customStringValue))

	hclFile := hclwrite.NewEmptyFile()

	hclFile.Body().AppendBlock(sig.Build())
	fmt.Println(string(hclFile.Bytes()))

	// Output:
	// my_block {
	//   attr1 = "basic_value"
	//   attr2 = local.my_local
	//   attr3 = var.my_var
	//   attr4 = data.my_data.my_property
	//   attr5 = "custom.my_var"
	//   attr6 = github_repository.res-id.name
	//   attr7 = [github_branch.res-id.branch, github_branch_protection.res-id.pattern]
	//
	//   attr8 = explicit_ident.foo
	//   attr9 = [explicit_ident_item.foo, explicit_ident_item.bar]
	//
	//   attr10 = custom.my_var
	// }
}

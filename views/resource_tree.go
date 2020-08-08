package views

import (
	"fmt"
	"strings"

	"github.com/wils0ns/tfscan/terraform"
)

type resourceTree struct{}

// Visit prints a tree of the state resources
func (v *resourceTree) Visit(module *terraform.Module, parent *terraform.Module) {
	prefix := "├─ "
	padding := ""

	if parent != nil {
		for i := 0; i < strings.Count(parent.Address, "module."); i++ {
			padding = "│  " + padding
		}
	}

	if module.Address != "" {
		fmt.Printf("%v%v%v:\n", padding, prefix, ColorModuleAddress(module.Address))
	}

	padding = ""
	for i := 0; i < strings.Count(module.Address, "module."); i++ {
		padding = "│  " + padding
	}

	for _, res := range module.Resources {
		address := ColorResourceAddress(res.Address)
		if module.Address != "" {
			address = fmt.Sprintf("%v.%v", ColorModuleAddress(module.Address), address)
		}

		if res.Index != "" {
			address = fmt.Sprintf("%v%v", address, ColorResourceIndex(fmt.Sprintf("[\"%v\"]", res.Index)))
		}

		fmt.Printf("%v%v%v\n", padding, prefix, address)

	}
}

// PrintResourceTree prints all resources within a state grouped by module
func PrintResourceTree(state *terraform.State) {
	for key, mod := range state.Values {
		fmt.Printf("%v:\n", ColorModuleAddress(key))
		mod.VisitModules(&resourceTree{}, nil)
	}
}

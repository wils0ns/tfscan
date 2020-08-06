package main

import (
	"fmt"
	"sort"
	"strings"
)

type resourcesTreeVisitor struct{}

// Visit prints a tree of the state resources
func (v *resourcesTreeVisitor) Visit(tfm *TerraformModule, parent *TerraformModule) {
	prefix := "├─ "
	padding := ""

	if parent != nil {
		for i := 0; i < strings.Count(parent.Address, "module."); i++ {
			padding = "│  " + padding
		}
	}

	if tfm.Address != "" {
		fmt.Printf("%v%v%v:\n", padding, prefix, tfm.Address)
	}

	padding = ""
	for i := 0; i < strings.Count(tfm.Address, "module."); i++ {
		padding = "│  " + padding
	}

	for _, res := range tfm.Resources {
		if res.Index != "" {
			fmt.Printf("%v%v%v[\"%v\"]\n", padding, prefix, blue(res.Address), res.Index)
			continue
		}
		fmt.Printf("%v%v%v\n", padding, prefix, blue(res.Address))

	}
}

// ResourceTypesVisitor stores the resource types within Terraform modyles
type ResourceTypesVisitor struct {
	typesMap map[string]struct{}
}

// NewResourceTypeVisitor initializes a resource
func NewResourceTypeVisitor() *ResourceTypesVisitor {
	return &ResourceTypesVisitor{typesMap: make(map[string]struct{})}
}

// Visit identify the unique resource types within Terraform modules
func (v *ResourceTypesVisitor) Visit(tfm *TerraformModule, parent *TerraformModule) {
	for _, res := range tfm.Resources {
		v.typesMap[res.Type] = struct{}{}
	}
	for _, cm := range tfm.ChildModules {
		v.Visit(cm, tfm)
	}
}

// Types returns a list of unique resource types
func (v *ResourceTypesVisitor) Types() ([]string, error) {
	if len(v.typesMap) > 0 {
		resTypes := make([]string, len(v.typesMap))
		i := 0
		for k := range v.typesMap {
			resTypes[i] = k
			i++
		}
		sort.Strings(resTypes)

		return resTypes, nil
	}

	return nil, fmt.Errorf("Not resources found")

}

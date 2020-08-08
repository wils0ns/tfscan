package terraform

import (
	"fmt"
	"sort"
)

// ResourceLookupVisitor stores the resource found by the given address
type ResourceLookupVisitor struct {
	Address  string
	Resource *Resource
}

// Visit searches for a resource that matches the Visitor's Address field
func (v *ResourceLookupVisitor) Visit(module, parent *Module) {
	if v.Resource != nil {
		return
	}

	for _, res := range module.Resources {
		address := res.Address
		if module.Address != "" {
			address = fmt.Sprintf("%v.%v", module.Address, address)
		}

		if res.Index != "" {
			address = fmt.Sprintf("%v[\"%v\"]", address, res.Index)
		}

		if v.Address == address {
			v.Resource = res
			return
		}
	}

	for _, cm := range module.ChildModules {
		v.Visit(cm, module)
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
func (v *ResourceTypesVisitor) Visit(m, parent *Module) {
	for _, res := range m.Resources {
		v.typesMap[res.Type] = struct{}{}
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

	return nil, fmt.Errorf("No resources found")

}

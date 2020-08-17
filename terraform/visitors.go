package terraform

import (
	"fmt"
	"log"
	"regexp"
	"sort"

	"github.com/Masterminds/semver"
)

// ResourceLookupVisitor stores the resource found by the given address
type ResourceLookupVisitor struct {
	TerraformVersion string
	AddressRegExp    string
	Resources        []*Resource
}

// Visit searches for resources that matches the Visitor's AddressRegExp field regular expression
// TODO: Design a way to capute errors from the visitor
func (v *ResourceLookupVisitor) Visit(module, parent *Module) {

	for _, res := range module.Resources {

		re, err := regexp.Compile(regexp.QuoteMeta(v.AddressRegExp))
		if err != nil {
			log.Println(err)
			return
		}

		address := res.Address
		if module.Address != "" {
			address = fmt.Sprintf("%v.%v", module.Address, address)
		}

		if res.Index != "" {
			address = fmt.Sprintf("%v[\"%v\"]", address, res.Index)
		}

		if re.MatchString(address) {
			version, err := semver.NewVersion(v.TerraformVersion)
			if err != nil {
				log.Println(err)
				return
			}
			c, err := semver.NewConstraint(">= 0.13.0")
			if err != nil {
				log.Println(err)
				return
			}
			if c.Check(version) {
				res.FullAddress = res.Address
			} else {
				res.FullAddress = address
			}
			v.Resources = append(v.Resources, res)
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

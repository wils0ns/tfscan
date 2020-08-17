package terraform

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Plan represents Terraform's plan
type Plan struct {
	FormatVersion    string             `json:"format_version"`
	TerraformVersion string             `json:"terraform_version"`
	Values           map[string]*Module `json:"planned_values"`
}

// NewPlan creates a new Plan
func NewPlan(r io.Reader) (*Plan, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	plan := &Plan{}
	err = json.Unmarshal(data, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

// ResourceLookup returns a slice of  resources that matches the given address regular expression
func (p *Plan) ResourceLookup(address string) ([]*Resource, error) {
	resLookup := ResourceLookupVisitor{AddressRegExp: address, TerraformVersion: p.TerraformVersion}
	for _, m := range p.Values {
		resLookup.Visit(m, nil)
	}

	if len(resLookup.Resources) == 0 {
		return nil, &ResourceNotFoundError{Address: address}
	}
	return resLookup.Resources, nil
}

// ResourceTypes returns a list of all the unique resources within the plan
func (p *Plan) ResourceTypes() ([]string, error) {
	v := NewResourceTypeVisitor()
	for _, m := range p.Values {
		m.VisitModules(v, nil)
	}

	return v.Types()
}

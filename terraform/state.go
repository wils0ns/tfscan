package terraform

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// State represents Terraform's state object
type State struct {
	FormatVersion    string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
	Values           map[string]*Module
}

// NewState creates a new TerraformState object
func NewState(r io.Reader) (*State, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	state := &State{}
	err = json.Unmarshal(data, state)
	if err != nil {
		return nil, err
	}

	return state, nil
}

// GetResourceByFullAddress returns a Resource that matches the given address
func (s *State) GetResourceByFullAddress(address string) (*Resource, error) {
	resLookup := ResourceLookupVisitor{Address: address}
	for _, m := range s.Values {
		resLookup.Visit(m, nil)
	}

	if resLookup.Resource == nil {
		return nil, &ResourceNotFoundError{Address: address}
	}
	return resLookup.Resource, nil
}

// ResourceTypes returns a list of all the unique resources within the state
func (s *State) ResourceTypes() ([]string, error) {
	v := NewResourceTypeVisitor()
	for _, m := range s.Values {
		m.VisitModules(v, nil)
	}

	return v.Types()
}

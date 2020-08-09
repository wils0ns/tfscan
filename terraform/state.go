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

// GetResourcesByFullAddress returns a slice of  resources that matches the given address regular expression
func (s *State) GetResourcesByFullAddress(address string) ([]*Resource, error) {
	resLookup := ResourceLookupVisitor{AddressRegExp: address}
	for _, m := range s.Values {
		resLookup.Visit(m, nil)
	}

	if len(resLookup.Resources) == 0 {
		return nil, &ResourceNotFoundError{Address: address}
	}
	return resLookup.Resources, nil
}

// ResourceTypes returns a list of all the unique resources within the state
func (s *State) ResourceTypes() ([]string, error) {
	v := NewResourceTypeVisitor()
	for _, m := range s.Values {
		m.VisitModules(v, nil)
	}

	return v.Types()
}

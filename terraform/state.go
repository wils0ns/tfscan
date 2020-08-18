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
	PlannedValues    map[string]*Module `json:"planned_values"`
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

// ActualValues returns the actual file values, depending if its a plan of state
func (s *State) ActualValues() map[string]*Module {
	values := s.Values
	if s.PlannedValues != nil {
		values = s.PlannedValues
	}
	return values
}

// ResourceLookup returns a slice of  resources that matches the given address regular expression
func (s *State) ResourceLookup(address string) ([]*Resource, error) {
	resLookup := ResourceLookupVisitor{AddressRegExp: address, TerraformVersion: s.TerraformVersion}

	for _, m := range s.ActualValues() {
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
	for _, m := range s.ActualValues() {
		m.VisitModules(v, nil)
	}

	return v.Types()
}

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// TerraformResource represents state resources
type TerraformResource struct {
	Address       string                 `json:"address"`
	Mode          string                 `json:"mode"`
	Type          string                 `json:"type"`
	Name          string                 `json:"name"`
	Index         string                 `json:"index"`
	ProviderName  string                 `json:"provider_name"`
	SchemaVersion int                    `json:"schema_version"`
	Values        map[string]interface{} `json:"values"`
}

// TerraformModule represents state top level modules
type TerraformModule struct {
	Resources    []*TerraformResource `json:"resources"`
	Address      string               `json:"address"`
	ChildModules []*TerraformModule   `json:"child_modules"`
}

// TerraformState represents Terraform's state object
type TerraformState struct {
	FormatVersion    string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
	Values           map[string]*TerraformModule
}

// TerraformModuleVisitor is the interface that wraps the Visit method
// The Visit Method takes the module to be visited and the its parent module
type TerraformModuleVisitor interface {
	Visit(*TerraformModule, *TerraformModule)
}

// TerraformResourceNotFoundError raised when resources are not found
type TerraformResourceNotFoundError struct {
	filter string
}

// NewTerraformState creates a new TerraformState object
func NewTerraformState(r io.Reader) (*TerraformState, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	state := &TerraformState{}
	err = json.Unmarshal(data, state)
	if err != nil {
		return nil, err
	}

	return state, nil
}

// VisitModules runs Visitor.Visit on the module and all its child modules
func (tfm *TerraformModule) VisitModules(visitor TerraformModuleVisitor, parent *TerraformModule) {
	visitor.Visit(tfm, parent)

	for _, childMod := range tfm.ChildModules {
		childMod.VisitModules(visitor, tfm)
	}
}

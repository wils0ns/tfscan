package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type TerraformChildModule struct {
	Resources    []*TerraformResource    `json:"resources"`
	Address      string                  `json:"address"`
	ChildModules []*TerraformChildModule `json:"child_modules"`
}

// TerraformResource struct
type TerraformResource struct {
	Address       string                 `json:"address"`
	Mode          string                 `json:"mode"`
	Type          string                 `json:"type"`
	Name          string                 `json:"name"`
	ProviderName  string                 `json:"provider_name"`
	SchemaVersion int                    `json:"schema_version"`
	Values        map[string]interface{} `json:"values"`
}

// TerraformModule struct
type TerraformModule struct {
	Resources    []*TerraformResource    `json:"resources"`
	ChildModules []*TerraformChildModule `json:"child_modules"`
}

// TerraformState struct
type TerraformState struct {
	FormatVersion    string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
	Values           map[string]*TerraformModule
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

func getChildResourceTypes(cm *TerraformChildModule) map[string]struct{} {
	typeMap := make(map[string]struct{})
	for _, res := range cm.Resources {
		typeMap[res.Type] = struct{}{}
	}

	for _, cm := range cm.ChildModules {
		for k := range getChildResourceTypes(cm) {
			typeMap[k] = struct{}{}
		}
	}
	return typeMap
}

// GetResourceTypes returns a slice of types of resources present in the state
func (tfs *TerraformState) GetResourceTypes() []string {
	typeMap := make(map[string]struct{})
	for _, res := range tfs.Values["root_module"].Resources {
		typeMap[res.Type] = struct{}{}
	}

	for _, cm := range tfs.Values["root_module"].ChildModules {
		childs := getChildResourceTypes(cm)
		for k := range childs {
			typeMap[k] = struct{}{}
		}
	}

	typeList := make([]string, len(typeMap))

	i := 0
	for k := range typeMap {
		typeList[i] = k
		i++
	}
	return typeList
}

package terraform

import "fmt"

// Resource represents state resources
type Resource struct {
	Address       string                 `json:"address"`
	Mode          string                 `json:"mode"`
	Type          string                 `json:"type"`
	Name          string                 `json:"name"`
	Index         string                 `json:"index"`
	ProviderName  string                 `json:"provider_name"`
	SchemaVersion int                    `json:"schema_version"`
	Values        map[string]interface{} `json:"values"`
}

// ResourceNotFoundError raised when resources are not found
type ResourceNotFoundError struct {
	Address string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Unable to find resource: %v", e.Address)
}

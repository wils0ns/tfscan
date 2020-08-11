package views

import (
	"encoding/json"
	"fmt"

	"github.com/wils0ns/tfscan/terraform"
)

// PrintResources prints the resource that matches the given address
func PrintResources(state *terraform.State, address string) error {
	resources, err := state.ResourceLookup(address)
	if err != nil {
		return err
	}

	resJSONBytes, err := json.MarshalIndent(resources, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(resJSONBytes))
	return nil
}

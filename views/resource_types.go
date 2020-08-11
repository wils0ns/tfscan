package views

import (
	"github.com/wils0ns/tfscan/terraform"
)

// PrintResourceTypes prints all resource types within a state
func PrintResourceTypes(state *terraform.State) error {
	resTypes, err := state.ResourceTypes()
	if err != nil {
		return err
	}

	return PrettyPrintJSON(resTypes)
}

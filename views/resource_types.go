package views

import (
	"fmt"
	"os"

	"github.com/wils0ns/tfscan/terraform"
)

// PrintResourceTypes prints all resource types within a state
func PrintResourceTypes(state *terraform.State) {
	resTypes, err := state.ResourceTypes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for _, item := range resTypes {
		fmt.Println(item)
	}
}

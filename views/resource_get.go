package views

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wils0ns/tfscan/terraform"
)

// PrintResource prints the resource that matches the given address
func PrintResource(state *terraform.State, address string) {
	res, err := state.GetResourceByFullAddress(address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	resource, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(resource))
}

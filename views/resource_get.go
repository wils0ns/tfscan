package views

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wils0ns/tfscan/terraform"
)

// PrintResources prints the resource that matches the given address
func PrintResources(state *terraform.State, address string) {
	resources, err := state.GetResourcesByFullAddress(address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	resJSONBytes, err := json.MarshalIndent(resources, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(resJSONBytes))
}

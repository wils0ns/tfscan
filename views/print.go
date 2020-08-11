package views

import (
	"encoding/json"
	"fmt"
	"os"
)

// PrintAndExitStdErr prints and error message and calls os.Exit(1)
func PrintAndExitStdErr(e error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", e)
	os.Exit(1)
}

// PrettyPrintJSON prints the JSON representation of the given object
func PrettyPrintJSON(i interface{}) error {
	iBytes, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(iBytes))
	return nil
}

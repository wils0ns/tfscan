package views

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kylelemons/godebug/diff"
)

// PrintDiff prints the different between two objects parsed to a indented JSON string
func PrintDiff(a, b interface{}) error {
	aBytes, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	bBytes, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	diffLines := strings.Split(diff.Diff(string(aBytes), string(bBytes)), "\n")

	for i, line := range diffLines {
		if strings.HasPrefix(line, "-") {
			diffLines[i] = ColorRed(line)
		} else if strings.HasPrefix(line, "+") {
			diffLines[i] = ColorGreen(line)
		}
	}

	fmt.Println(strings.Join(diffLines, "\n"))
	return nil
}

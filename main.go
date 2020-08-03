package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/fatih/color"
)

// Command line
var flagNoColor = flag.Bool("no-color", false, "Disable color output")
var flagJSONFile = flag.String("json", "", "Path to a state or plan JSON file")
var flagResTypes = flag.Bool("types", false, "Only list the type of resources present in the state")
var flagResTree = flag.Bool("tree", false, "Print resources grouped by module")

// ListResources lists all resources within a state
func ListResources(tfs *TerraformState) {
	var resources []string
	cyan := color.New(color.FgCyan).SprintFunc()
	for _, res := range tfs.Values["root_module"].Resources {
		fmt.Println(cyan(res.Address), res.Values["id"])
		// fmt.Println(res)
		resources = append(resources, res.Name)
	}
}

// ProcessState process command requests around the state
func ProcessState(r io.Reader) {
	state, err := NewTerraformState(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if *flagResTypes {
		resTypes := state.GetResourceTypes()
		sort.Strings(resTypes)
		for _, resType := range resTypes {
			fmt.Println(resType)
		}
		os.Exit(0)
	}

	if *flagResTree {
		// WIP
	}

	ListResources(state)
	os.Exit(0)
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Println("Example 01: tfscan -json state.json")
		fmt.Println("Example 02: terraform show -json | tfscan")
		fmt.Print("\nCommand line flags:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	if *flagNoColor {
		color.NoColor = true
	}

	// Read state from STDIN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		ProcessState(os.Stdin)
	}

	// Read state from file
	if *flagJSONFile != "" {
		jsonFile, err := os.Open(*flagJSONFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer jsonFile.Close()

		ProcessState(jsonFile)
	}

	if len(os.Args) < 2 {
		flag.Usage()
	}

}

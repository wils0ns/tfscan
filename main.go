package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// Command line
var flagNoColor = flag.Bool("no-color", false, "Disable color output")
var flagJSONFile = flag.String("json", "", "Path to a state or plan JSON file")
var flagResTypes = flag.Bool("types", false, "Only list the type of resources present in the state")
var flagResTree = flag.Bool("tree", false, "Print resources grouped by module")

func printResourceTree(tfs *TerraformState) {
	for key, mod := range tfs.Values {
		fmt.Printf("%v:\n", key)
		mod.VisitModules(&resourcesTreeVisitor{}, nil)
	}
}

func printResourceTypes(tfs *TerraformState) {
	for _, mod := range tfs.Values {
		v := NewResourceTypeVisitor()
		mod.VisitModules(v, nil)
		resTypes, err := v.Types()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		for _, item := range resTypes {
			fmt.Println(item)
		}
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
		printResourceTypes(state)
		os.Exit(0)
	}

	printResourceTree(state)
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

	// Print tool usage if no arguments have been passed
	if len(os.Args) < 2 {
		flag.Usage()
	}

}

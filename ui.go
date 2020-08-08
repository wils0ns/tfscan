package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/wils0ns/tfscan/terraform"
	"github.com/wils0ns/tfscan/views"
)

// Command line flags
var flagNoColor = flag.Bool("no-color", false, "Disable color output")
var flagJSONFile = flag.String("json", "", "Path to a state or plan JSON file")
var flagResTypes = flag.Bool("types", false, "Only list the type of resources present in the state")
var flagResGet = flag.String("get", "", "Shows the content of the resource that matches the given full address")

// var flagResFind = flag.String("find", "", "Shows the content of all resources where the full address matches the given regular expression")
// var flagStateDiff = flag.String("diff", "", "Loads a second state or plan JSON file and shows the difference between them")

type example struct {
	Descrition string
	Command    string
}

func printExamples() {
	fmt.Println("Examples:")
	examples := []example{
		{Descrition: "Read state from file path", Command: "tfscan -json state.json"},
		{Descrition: "Read state from stdin", Command: "terraform show -json | tfscan"},
		{Descrition: "Get resource content", Command: "terraform show -json | tfscan -get module.project.google_project.default"},
		{Descrition: "Get indexed resource content", Command: "terraform show -json | tfscan -get google_project_service.default[\"iam.googleapis.com\"]"},
	}
	for _, e := range examples {
		fmt.Printf("  %v:\n\t%v\n\n", e.Descrition, e.Command)
	}
}

func initUI() {

	color.Unset()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s usage:\n\n\n", os.Args[0])
		printExamples()
		fmt.Print("\nCommand line flags:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func loadState() *terraform.State {

	reader := os.Stdin
	stat, _ := reader.Stat()

	// If nothing to read from stdin, expecting -json flag
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		jsonFile, err := os.Open(*flagJSONFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer jsonFile.Close()
		reader = jsonFile
	}

	state, err := terraform.NewState(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	return state
}

func parseCommandLine(args []string) {
	if args == nil {
		args = os.Args[1:]
	}
	flag.CommandLine.Parse(args)

	if *flagNoColor {
		color.NoColor = true
	}

	state := loadState()

	if *flagResTypes {
		views.PrintResourceTypes(state)
		os.Exit(0)
	}

	if *flagResGet != "" {
		views.PrintResource(state, *flagResGet)
		os.Exit(0)
	}

	views.PrintResourceTree(state)
	os.Exit(0)

	// Print tool usage if no arguments have been passed
	if len(args) < 1 {
		flag.Usage()
	}
}

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
var flagDiff = flag.String("diff", "", "Loads a second state or plan JSON file and shows the difference between them")
var flagResTypes = flag.Bool("types", false, "Only list the type of resources present in the state")
var flagResGet = flag.String("get", "", "Shows the content of the resource that matches the given full address")

// var flagResFind = flag.String("find", "", "Shows the content of all resources where the full address matches the given regular expression")

func printExamples() {
	fmt.Println("Examples:")
	examples := []struct {
		Descrition string
		Command    string
	}{
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

func loadStateFromJSON(name string) (*terraform.State, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("%v(%v)", err, *flagJSONFile)
	}
	defer jsonFile.Close()
	state, err := terraform.NewState(jsonFile)
	if err != nil {
		return nil, err

	}
	return state, nil
}

func loadState() (*terraform.State, error) {

	reader := os.Stdin
	stat, _ := reader.Stat()

	// If something to read from stdin, expecting -json flag
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		state, err := terraform.NewState(reader)
		if err != nil {
			return nil, err
		}
		return state, nil

	} else if *flagJSONFile != "" {
		state, err := loadStateFromJSON(*flagJSONFile)
		if err != nil {
			return nil, fmt.Errorf("%v(%v)", err, *flagJSONFile)
		}
		return state, nil

	} else {
		return nil, fmt.Errorf("error: A valid state file must be specified")
	}
}

func parseCommandLine(args []string) {
	if args == nil {
		args = os.Args[1:]
	}
	flag.CommandLine.Parse(args)

	if *flagNoColor {
		color.NoColor = true
	}

	state, err := loadState()
	if err != nil {
		views.PrintAndExitStdErr(err)
	}

	// TODO: make parser DRYer

	if *flagResTypes {
		if *flagDiff != "" {
			otherState, err := loadStateFromJSON(*flagDiff)
			if err != nil {
				views.PrintAndExitStdErr(err)
			}
			resA, err := state.ResourceTypes()
			if err != nil {
				views.PrintAndExitStdErr(err)
			}

			resB, err := otherState.ResourceTypes()
			if err != nil {
				views.PrintAndExitStdErr(err)
			}

			err = views.PrintDiff(resA, resB)
			if err != nil {
				views.PrintAndExitStdErr(err)
			}
		} else {
			err = views.PrintResourceTypes(state)
			if err != nil {
				views.PrintAndExitStdErr(err)
			}
		}

		os.Exit(0)
	}

	if *flagResGet != "" {
		if *flagDiff != "" {
			otherState, err := loadStateFromJSON(*flagDiff)
			if err != nil {
				views.PrintAndExitStdErr(err)
			}
			resA, err := state.ResourceLookup(*flagResGet)
			if err != nil {
				err = fmt.Errorf("%v: %v", *flagJSONFile, err)
				views.PrintAndExitStdErr(err)
			}

			resB, err := otherState.ResourceLookup(*flagResGet)
			if err != nil {
				err = fmt.Errorf("%v: %v", *flagDiff, err)
				views.PrintAndExitStdErr(err)
			}

			err = views.PrintDiff(resA, resB)
			if err != nil {
				views.PrintAndExitStdErr(err)
			}
		} else {
			err = views.PrintResources(state, *flagResGet)
			if err != nil {
				views.PrintAndExitStdErr(err)
			}
		}

		os.Exit(0)
	}

	views.PrintResourceTree(state)
	os.Exit(0)

	// Print tool usage if no arguments have been passed
	if len(args) < 1 {
		flag.Usage()
	}
}

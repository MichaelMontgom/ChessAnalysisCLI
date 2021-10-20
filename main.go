package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var openingMoveCommand string
	flag.StringVar(&openingMoveCommand, "opening", "", "Pass user's chess.com profile name. Default is \"\"")

	flag.Parse()

	// see if enough arguments have been passed
	if len(os.Args) < 2 {
		fmt.Println("subcommand required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-opening":
		// fmt.Print(getOpeningMovePreference(openingMoveCommand))
		fmt.Print(getMonths())
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

}

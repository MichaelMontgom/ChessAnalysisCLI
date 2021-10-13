package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	// fmt.Print(getPlayerProfile("DarthBraves"))

	openingMoveCommand := flag.NewFlagSet("opening", flag.ExitOnError)

	openingGames := openingMoveCommand.String("games", "", "Pass user's chess.com name to get their past games opening moves")

	// see if enough arguments have been passed
	if len(os.Args) < 2 {
		log.Println("subcommand required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "opening":
		openingMoveCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if openingMoveCommand.Parsed() {
		if *openingGames == "" {
			openingMoveCommand.PrintDefaults()
			os.Exit(1)
		}
		// logic for -games

		getOpeningMovePreference(*openingGames)

	}

}

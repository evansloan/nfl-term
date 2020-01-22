package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/evansloan/nfl-term/api"
)

var (
	version string
)

func main() {

	var games []*api.Game

	if len(os.Args) > 1 {
		arg := strings.ToLower(os.Args[1])
		if arg == "--version" || arg == "-v" {
			fmt.Println("nfl-term version: ", version)
			os.Exit(3)
		} else if arg == "--games" || arg == "-g" {
			if len(os.Args) < 3 {
				fmt.Println("--games requires at least 1 game ID\nex. nfl-term --games 2015091000 2015091301")
				os.Exit(3)
			}
			games = api.GetGames(os.Args[2:len(os.Args)])
		} else {
			fmt.Println("Invalid argument.")
		}
	} else {
		games = api.Games()
	}

	app := NewApp(games)
	app.RunApp()
}

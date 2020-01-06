package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	version string
)

func main() {

	if len(os.Args) > 1 {
		arg := os.Args[1]
		if strings.ToLower(arg) == "version" {
			fmt.Println("nfl-term version: ", version)
			os.Exit(3)
		}
	}

	app := NewApp()
	app.RunApp()
}

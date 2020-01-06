package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	version string
	commit  string
)

func main() {

	if len(os.Args) > 1 {
		arg := os.Args[1]
		if strings.ToLower(arg) == "version" {
			fmt.Println("Version: ", version)
			fmt.Println("Commit: ", commit)
			os.Exit(3)
		}
	}

	app := NewApp()
	app.RunApp()
}

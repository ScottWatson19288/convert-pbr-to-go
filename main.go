package main

import (
	"fmt"

	api "github.com/ScottWatson19288/convert-pbr-to-go/api"
)

var appState api.Options

func main() {
	// Parse arguments

	// Set Options
	err := appState.Init()
	if err != nil {
		fmt.Println("failed to init options")
	} else {
		// Parse Scene File
		err = api.ParseFile("")
		if err != nil {
			fmt.Println("failed to load scene file")
		}
	}
	// Clean Up and Exit
	api.Cleanup()
}

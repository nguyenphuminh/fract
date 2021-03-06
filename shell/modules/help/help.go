package help

import (
	"fmt"

	"../../../utilities/cli"
)

// Process Process command in module.
// command Command to process.
func Process(command string) {
	if command != "" {
		fmt.Println("This module can only be used!")
		return
	}
	cli.PrintMapAsTable(map[string]string{
		"make":    "Interprete Fract code.",
		"version": "Show version.",
		"help":    "Show help.",
		"exit":    "Exit.",
	})
}

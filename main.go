package main

import (
	"bufio"
	"jash-go/helpers"
	"os"
)

func main() {
	// Commands
	/* Command map (CMDS) is declared externally in consts.go */

	// Aliases
	/* Alias map (ALIASES) is declared externally in consts.go */

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Get current command
		helpers.PrintPromptStr()
		scanner.Scan()
		cmd := helpers.ParseInput(scanner.Text())

		// Check if command is alias
		if alias, ok := helpers.ALIASES[cmd.Cmd]; ok {
			// Reparse in case alias had arguments itself
			cmd = helpers.ParseInput(alias + " " + cmd.ArgStr())
		}

		// Find & execute command
		if fnc, ok := helpers.CMDS[cmd.Cmd]; ok {
			fnc.Function(cmd.Argv)
		} else {
			println(cmd.Cmd + ": command not found")
		}
	}
}

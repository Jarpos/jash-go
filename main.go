package main

import (
	"bufio"
	"fmt"
	"jash-go/helpers"
	"os"
)

func main() {
	// Commands
	/* Command map (CMDS) is declared externally in consts.go */
	initCommands()

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
			err := fnc.Function(cmd.Argv)
			if err != nil {
				// Handle error
				// TODO: Make errors actually do something
				fmt.Println("Error:", err.Error())
			}
		} else {
			fmt.Println(cmd.Cmd + ": command not found")
		}
	}
}

func initCommands() {
	// Add help command here to avoid InvalidInitCycle error
	helpers.CMDS["help"] = helpers.BuiltinCommand{
		Description: "Prints this help", Function: helpers.Help,
	}
	helpers.CMDS["alias"] = helpers.BuiltinCommand{
		Description: "Print all aliases or add new one", Function: helpers.Alias,
	}
}

package main

import (
	"bufio"
	"jash-go/builtins"
	"jash-go/helpers"
	"os"
)

func main() {
	cmds := map[string]func([]string){
		"ls": builtins.ListSubdirectories,
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Get current command
		helpers.PrintPromptStr()
		scanner.Scan()
		cmd := helpers.ParseInput(scanner.Text())

		// Find & execute command
		if fnc, ok := cmds[cmd.Cmd]; ok {
			fnc(cmd.Argv)
		} else {
			println(cmd.Cmd + ": command not found")
		}
	}
}

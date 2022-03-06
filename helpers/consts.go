package helpers

import (
	"jash-go/builtins"
)

type BuiltinCommand struct {
	Description string
	Function    func([]string)
}

// Builtin commands
var CMDS = map[string]BuiltinCommand{
	"ls":  {"Changes current directory" /*****/, builtins.ListSubdirectories},
	"cd":  {"Lists directory contents" /******/, builtins.ChangeDirectory},
	"cat": {"Print file content to stdout" /**/, builtins.ConcatenateAndPrint},
}

// Command aliases
var ALIASES = map[string]string{
	"l": "ls",
}

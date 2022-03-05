package helpers

import (
	"jash-go/builtins"
)

// Builtin commands
var CMDS = map[string]func([]string){
	"ls": builtins.ListSubdirectories,
}

// Command aliases
var ALIASES = map[string]string{
	"l": "ls",
}

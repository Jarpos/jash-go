package helpers

import (
	"fmt"
	"jash-go/builtins"
)

type BuiltinCommand struct {
	Description string
	Function    func([]string)
}

// Builtin commands
var CMDS = map[string]BuiltinCommand{
	"ls":     {"Changes current directory" /**********/, builtins.ListSubdirectories},
	"cd":     {"Lists directory contents" /***********/, builtins.ChangeDirectory},
	"cat":    {"Prints contents of file to stdout" /**/, builtins.ConcatenateAndPrint},
	"mkdir":  {"Creates new directory" /**************/, builtins.MakeDirectory},
	"mkfile": {"Creates new, empty file" /************/, builtins.MakeFile},
	//"help": {"Prints this help", /******************/, helpers.Help}, <-- added in main.go
}

// Command aliases
var ALIASES = map[string]string{
	"l": "ls",
}

func Help(argv []string) {
	for key, val := range CMDS {
		fmt.Printf("%s:\t%s\n", key, val.Description)
	}
}

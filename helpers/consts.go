package helpers

import (
	"fmt"
	"jash-go/builtins"
	"strings"
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
	"tree":   {"Shows directory tree" /***************/, builtins.Tree},
	"exit":   {"Exits the shell" /********************/, builtins.Exit},
	//"help":  {"Prints this help", /****************\/, helpers.Help}, <-- added in main.go
	//"alias": {"Prints all aliases", /**************\/, helpers.Aliases}, <-- added in main.go
}

// Command aliases
var ALIASES = map[string]string{
	"l":  "ls",
	"..": "cd ..",
}

func Help(argv []string) {
	for key, val := range CMDS {
		fmt.Printf("%7s %s\n", key+":", val.Description)
	}
}

func Alias(argv []string) {
	if len(argv) == 0 {
		for key, val := range ALIASES {
			fmt.Printf("alias %s='%s'\n", key, val)
		}
	} else {
		ALIASES[argv[0]] = strings.Join(argv[1:], " ")
	}
}

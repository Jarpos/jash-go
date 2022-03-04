package builtins

import (
	"fmt"
	"os"
)

func ListSubdirectories(argv []string) {
	p := "."
	if len(argv) != 0 {
		p = argv[0]
	}

	files, err := os.ReadDir(p)
	if err != nil {
		fmt.Println("Error path: " + p + " does not exist")
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

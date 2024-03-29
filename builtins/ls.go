package builtins

import (
	"fmt"
	"os"
)

func ListSubdirectories(argv []string) error {
	p := "."
	if len(argv) != 0 {
		p = argv[0]
	}

	files, err := os.ReadDir(p)
	if err != nil {
		fmt.Println("Error on: " + err.Error())
		return nil
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}

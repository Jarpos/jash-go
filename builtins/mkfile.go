package builtins

import (
	"fmt"
	"os"
)

func MakeFile(argv []string) {
	if len(argv) == 0 {
		fmt.Println("Error: Please specify a file name")
		return
	}

	name := argv[0]
	file, err := os.Stat(name)
	if err == nil && !file.IsDir() {
		fmt.Println("Error: " + file.Name() + " exists already")
		return
	}

	_, err = os.Create(name)
	if err != nil {
		fmt.Println("Error on: " + err.Error())
	}
}

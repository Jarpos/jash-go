package builtins

import (
	"fmt"
	"os"
)

func MakeDirectory(argv []string) {
	if len(argv) == 0 {
		fmt.Println("Error: Please specify a file")
	}

	name := argv[0]
	file, err := os.Stat(name)
	if err == nil && file.IsDir() {
		fmt.Println("Error: " + file.Name() + " exists already and is a directory")
		return
	}

	err = os.Mkdir(name, 0777)
	if err != nil {
		fmt.Println("Error on: " + err.Error())
	}
}

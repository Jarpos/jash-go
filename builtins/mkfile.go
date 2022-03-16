package builtins

import (
	"fmt"
	"os"
)

func MakeFile(argv []string) error {
	if len(argv) == 0 {
		fmt.Println("Error: Please specify a file name")
		return nil
	}

	name := argv[0]
	file, err := os.Stat(name)
	if err == nil && !file.IsDir() {
		fmt.Println("Error: " + file.Name() + " exists already")
		return nil
	}

	_, err = os.Create(name)
	if err != nil {
		fmt.Println("Error on: " + err.Error())
	}
	return nil
}

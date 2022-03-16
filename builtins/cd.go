package builtins

import (
	"fmt"
	"os"
)

func ChangeDirectory(argv []string) error {
	path, _ := os.UserHomeDir()
	if len(argv) >= 1 {
		path = argv[0]
	}

	file, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error on: " + err.Error())
		return nil
	}
	if !file.IsDir() {
		fmt.Println("Error: " + file.Name() + " is not a directory")
		return nil
	}

	os.Chdir(path)
	return nil
}

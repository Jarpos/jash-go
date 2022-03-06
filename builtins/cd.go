package builtins

import (
	"fmt"
	"os"
)

func ChangeDirectory(argv []string) {
	path, _ := os.UserHomeDir()
	if len(argv) >= 1 {
		path = argv[0]
	}

	file, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error on: " + err.Error())
		return
	}
	if !file.IsDir() {
		fmt.Println("Error: " + file.Name() + " is not a directory")
		return
	}

	os.Chdir(path)
}

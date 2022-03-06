package builtins

import (
	"fmt"
	"os"
)

func ConcatenateAndPrint(argv []string) {
	if len(argv) == 0 {
		fmt.Println("Error: Please specify a file")
	}

	path := argv[0]
	file, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error on: " + err.Error())
		return
	}
	if file.IsDir() {
		fmt.Println("Error: " + file.Name() + " is not a file")
		return
	}

	content, _ := os.ReadFile(path)
	fmt.Println(string(content))
}

package builtins

import (
	"fmt"
	"os"
)

func Tree(argv []string) {
	path := "."
	if len(argv) != 0 {
		path = argv[0]
	}

	traverse(path, "")
}

func traverse(dir string, level string) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(level + file.Name())
			traverse(dir+"/"+file.Name(), level+"    ")
		} else {
			fmt.Println(level + file.Name())
		}
	}
}

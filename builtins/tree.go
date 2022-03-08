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

	_, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error on: " + err.Error())
		return
	}

	fmt.Println(path)
	traverse(path, 1)
}

func traverse(dir string, level int) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(getOffset(level) + file.Name())
			traverse(dir+"/"+file.Name(), level+1)
		} else {
			fmt.Println(getOffset(level) + file.Name())
		}
	}
}

func getOffset(level int) string {
	s := ""
	for i := 0; i < level; i++ {
		s += "    "
	}
	return s
}

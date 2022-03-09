package builtins

import (
	"fmt"
	"io/fs"
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

	onTraverse := func(f fs.DirEntry, level int) {
		if f.IsDir() {
			fmt.Println(getPadding(level) + f.Name())
		} else {
			fmt.Println(getPadding(level) + f.Name())
		}
	}

	fmt.Println(path)
	traverse(path, 1, onTraverse)
}

func traverse(dir string, level int, onTraverse func(fs.DirEntry, int)) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		onTraverse(file, level)
		if file.IsDir() {
			traverse(dir+"/"+file.Name(), level+1, onTraverse)
		}
	}
}

func getPadding(level int) string {
	s := ""
	for i := 0; i < level; i++ {
		s += "    "
	}
	return s
}

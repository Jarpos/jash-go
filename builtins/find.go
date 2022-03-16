package builtins

import (
	"fmt"
	"io/fs"
	"os"
)

func Find(argv []string) {
	if len(argv) != 2 {
		fmt.Println("Usage: find [filename] [startfolder]")
		return
	}

	var results []string
	fnc := func(f fs.DirEntry, path string) {
		if f.Name() == argv[0] {
			results = append(results, path)
		}
	}
	traverseFind(argv[1], fnc)
	fmt.Println("Results:")
	for _, result := range results {
		fmt.Println(result)
	}
}

func traverseFind(dir string, onTraverse func(fs.DirEntry, string)) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		onTraverse(file, dir+"/"+file.Name())
		if file.IsDir() {
			traverseFind(dir+"/"+file.Name(), onTraverse)
		}
	}
}

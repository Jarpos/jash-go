package builtins

import (
	"errors"
	"os"
)

func MakeFile(argv []string) error {
	if len(argv) == 0 {
		return errors.New("no file name specified")
	}

	name := argv[0]
	file, err := os.Stat(name)
	if err == nil && !file.IsDir() {
		return errors.New(file.Name() + " exists already")
	}

	_, err = os.Create(name)
	return err
}

package builtins

import (
	"os"
)

func Exit(argv []string) error {
	os.Exit(0)
	return nil
}

package builtins

import (
	"os"
)

func Exit(argv []string) {
	os.Exit(0)
}

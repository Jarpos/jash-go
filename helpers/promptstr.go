package helpers

import (
	"fmt"
	"os"
)

func PrintPromptStr() {
	path, _ := os.Getwd()
	fmt.Print(path + "$ ")
}

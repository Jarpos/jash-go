package helpers

import (
	"regexp"
	"strings"
)

type Command struct {
	Cmd  string
	Argv []string
}

func ParseInput(i string) Command {
	spaces := regexp.MustCompile(`\s+`)
	i = spaces.ReplaceAllString(i, " ")
	i = strings.TrimSpace(i)
	tokens := strings.Split(i, " ")

	var c = Command{
		Cmd:  tokens[0],
		Argv: tokens[1:],
	}

	return c
}

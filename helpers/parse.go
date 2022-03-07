package helpers

import (
	"regexp"
	"strings"
)

type Command struct {
	Cmd  string
	Argv []string
}

func (c Command) ArgStr() string {
	return strings.Join(c.Argv, " ")
}

func ParseInput(i string) Command {
	reg := regexp.MustCompile(`\s+`)
	i = reg.ReplaceAllString(i, " ")
	i = strings.TrimSpace(i)
	tokens := strings.Split(i, " ")

	return Command{
		Cmd:  tokens[0],
		Argv: tokens[1:],
	}
}

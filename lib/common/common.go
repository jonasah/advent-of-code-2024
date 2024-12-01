package common

import (
	"os"
	"strings"
)

func GetInput(day int) string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(content))
}

func GetLines(input string) []string {
	return strings.Split(input, "\n")
}

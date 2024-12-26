package day25

import (
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
)

const (
	filled    = '#'
	filledRow = "#####"
)

func Part1(input string) int {
	locks, keys := parseInput(input)

	count := 0
	for _, lock := range locks {
	loop:
		for _, key := range keys {
			for c := range lock {
				if lock[c]+key[c] > 5 {
					continue loop
				}
			}

			count++
		}
	}

	return count
}

func parseInput(input string) ([][]int, [][]int) {
	schematics := strings.Split(input, "\n\n")

	var locks [][]int
	var keys [][]int

	for _, s := range schematics {
		rows := common.GetLines(s)

		heights := make([]int, len(rows[0]))
		for _, r := range rows[1 : len(rows)-1] {
			for c, x := range r {
				if x == filled {
					heights[c]++
				}
			}
		}

		if rows[0] == filledRow {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	return locks, keys
}

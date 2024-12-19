package day19

import (
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

func Part1(input string) int {
	patterns, designs := parseInput(input)

	memo := make(map[string]int)

	sum := 0
	for _, d := range designs {
		if count(patterns, d, memo) > 0 {
			sum++
		}
	}

	return sum
}

func Part2(input string) int {
	patterns, designs := parseInput(input)

	memo := make(map[string]int)

	sum := 0
	for _, d := range designs {
		sum += count(patterns, d, memo)
	}

	return sum
}

func parseInput(input string) ([]string, []string) {
	lines := common.GetLines(input)

	patterns := strings.FieldsFunc(lines[0], func(r rune) bool { return r == ' ' || r == ',' })
	designs := lines[2:]

	return patterns, designs
}

func count(patterns []string, design string, memo map[string]int) int {
	if design == "" {
		return 1
	}

	c, inMemo := memo[design]
	if inMemo {
		return c
	}

	idx := slicelib.IndexAllFunc(patterns, func(p string) bool { return strings.HasPrefix(design, p) })
	for _, i := range idx {
		memo[design] += count(patterns, strings.TrimPrefix(design, patterns[i]), memo)
	}

	return memo[design]
}

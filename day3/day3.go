package day3

import (
	"regexp"
	"strconv"

	"github.com/jonasah/advent-of-code-2024/lib/common"
)

func Part1(input string) int {
	lines := common.GetLines(input)

	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0
	for _, l := range lines {
		matches := r.FindAllStringSubmatch(l, -1)

		for _, m := range matches {
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])
			sum += x * y
		}
	}

	return sum
}

func Part2(input string) int {
	lines := common.GetLines(input)

	r := regexp.MustCompile(`(?:mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)

	sum := 0
	enabled := true
	for _, l := range lines {
		matches := r.FindAllStringSubmatch(l, -1)

		for _, m := range matches {
			if m[0] == "do()" {
				enabled = true
				continue
			}

			if m[0] == "don't()" {
				enabled = false
				continue
			}

			if !enabled {
				continue
			}

			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])
			sum += x * y
		}
	}

	return sum
}

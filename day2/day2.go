package day2

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

func Part1(input string) int {
	reports := common.GetLines(input)

	safe := 0
	for _, r := range reports {
		levels := slicelib.Atoi(strings.Fields(r))

		if isSafe(levels) {
			safe++
		}
	}

	return safe
}

func Part2(input string) int {
	reports := common.GetLines(input)

	safe := 0
	for _, r := range reports {
		levels := slicelib.Atoi(strings.Fields(r))

		if isSafe(levels) {
			safe++
			continue
		}

		for i := range levels {
			newLevels := slices.Delete(slices.Clone(levels), i, i+1)

			if isSafe(newLevels) {
				safe++
				break
			}
		}
	}

	return safe
}

func isSafe(levels []int) bool {
	if levels[0] > levels[len(levels)-1] {
		slices.Reverse(levels)
	}

	d := diff(levels)

	return !slices.ContainsFunc(d, func(d int) bool {
		return d < 1 || d > 3
	})
}

func diff(a []int) []int {
	d := make([]int, 0, len(a)-1)
	for i := range a[1:] {
		d = append(d, a[i+1]-a[i])
	}
	return d
}

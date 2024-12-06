package day5

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/sliceconv"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

func Part1(input string) int {
	inputParts := strings.Split(input, "\n\n")

	rules := parseRules(inputParts[0])
	updates := parseUpdates(inputParts[1])

	sum := 0
	for _, u := range updates {
		if slices.IsSortedFunc(u, compareFn(rules)) {
			sum += u[len(u)/2]
		}
	}

	return sum
}

func Part2(input string) int {
	inputParts := strings.Split(input, "\n\n")

	rules := parseRules(inputParts[0])
	updates := parseUpdates(inputParts[1])

	sum := 0
	for _, u := range updates {
		before := slices.Clone(u)
		slices.SortStableFunc(u, compareFn(rules))

		if slices.Compare(before, u) != 0 {
			sum += u[len(u)/2]
		}
	}

	return sum
}

func parseRules(rulesInput string) map[int][]int {
	lines := common.GetLines(rulesInput)

	rules := map[int][]int{}
	for _, l := range lines {
		v := strings.Split(l, "|")
		p1, _ := strconv.Atoi(v[0])
		p2, _ := strconv.Atoi(v[1])
		rules[p1] = append(rules[p1], p2)
	}

	return rules
}

func parseUpdates(updatesInput string) [][]int {
	lines := common.GetLines(updatesInput)
	return slicelib.Map(lines, func(l string) []int { return sliceconv.Atoi(strings.Split(l, ",")) })
}

func compareFn(rules map[int][]int) func(a, b int) int {
	return func(a, b int) int {
		aFirst := slices.Contains(rules[a], b)
		bFirst := slices.Contains(rules[b], a)

		if !aFirst && !bFirst {
			return 0
		}

		if aFirst {
			return -1
		}

		return 1
	}
}

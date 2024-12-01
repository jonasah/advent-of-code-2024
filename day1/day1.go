package day1

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
)

func Part1(input string) int {
	lines := common.GetLines(input)

	var left, right []int
	for _, l := range lines {
		cols := strings.Fields(l)

		lv, _ := strconv.Atoi(cols[0])
		left = append(left, lv)

		rv, _ := strconv.Atoi(cols[1])
		right = append(right, rv)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i, lv := range left {
		rv := right[i]
		sum += mathlib.Abs(lv - rv)
	}

	return sum
}

func Part2(input string) int {
	lines := common.GetLines(input)

	var left []int
	right := make(map[int]int)
	for _, l := range lines {
		cols := strings.Fields(l)

		lv, _ := strconv.Atoi(cols[0])
		left = append(left, lv)

		rv, _ := strconv.Atoi(cols[1])
		right[rv]++
	}

	sum := 0
	for _, lv := range left {
		sum += lv * right[lv]
	}

	return sum
}

package day7

import (
	"math"
	"strings"
	"unicode"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

type equation struct {
	testValue int
	numbers   []int
}

func Part1(input string) int {
	return solve(input, 1)
}

func Part2(input string) int {
	return solve(input, 2)
}

func solve(input string, part int) int {
	equations := parseInput(input)

	sum := 0
	for _, e := range equations {
		solved := calc(e.testValue, e.numbers[0], e.numbers[1:], part)
		if solved {
			sum += e.testValue
		}
	}

	return sum
}

func parseInput(input string) []equation {
	lines := common.GetLines(input)

	var equations []equation
	for _, l := range lines {
		v := slicelib.Atoi(strings.FieldsFunc(l, func(r rune) bool { return !unicode.IsDigit(r) }))
		equations = append(equations, equation{v[0], v[1:]})
	}

	return equations
}

func calc(testValue, result int, numbers []int, part int) bool {
	if result > testValue {
		return false
	}

	if len(numbers) == 0 {
		return result == testValue
	}

	if calc(testValue, result+numbers[0], numbers[1:], part) {
		return true
	}

	if calc(testValue, result*numbers[0], numbers[1:], part) {
		return true
	}

	if part == 1 {
		return false
	}

	shift := int(math.Floor(math.Log10(float64(numbers[0])))) + 1
	concat := result*int(math.Pow10(shift)) + numbers[0]
	return calc(testValue, concat, numbers[1:], part)
}

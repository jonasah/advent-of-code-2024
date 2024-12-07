package day7

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 3749, Part1(testInput))
	require.Equal(t, 303876485655, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 11387, Part2(testInput))
	require.Equal(t, 146111650210682, Part2(realInput))
}

var realInput = common.GetInput(7)

const testInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

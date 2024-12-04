package day4

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 18, Part1(testInput))
	require.Equal(t, 2562, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 9, Part2(testInput))
	require.Equal(t, 1902, Part2(realInput))
}

var realInput = common.GetInput(4)

const testInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

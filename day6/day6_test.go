package day6

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 41, Part1(testInput))
	require.Equal(t, 4696, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 6, Part2(testInput))
	require.Equal(t, 1443, Part2(realInput))
}

var realInput = common.GetInput(6)

const testInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

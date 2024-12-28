package day20

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 8, Part1(testInput, 12))
	require.Equal(t, 1369, Part1(realInput, 100))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 41, Part2(testInput, 70))
	require.Equal(t, 979012, Part2(realInput, 100))
}

var realInput = common.GetInput(20)

const testInput = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

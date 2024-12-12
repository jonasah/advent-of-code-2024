package day12

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 140, Part1(testInput1))
	require.Equal(t, 772, Part1(testInput2))
	require.Equal(t, 1930, Part1(testInput3))
	require.Equal(t, 1387004, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 80, Part2(testInput1))
	require.Equal(t, 436, Part2(testInput2))
	require.Equal(t, 236, Part2(testInput4))
	require.Equal(t, 368, Part2(testInput5))
	require.Equal(t, 1206, Part2(testInput3))
	require.Equal(t, 844198, Part2(realInput))
}

var realInput = common.GetInput(12)

const testInput1 = `AAAA
BBCD
BBCC
EEEC`

const testInput2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

const testInput3 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

const testInput4 = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

const testInput5 = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

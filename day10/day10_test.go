package day10

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 36, Part1(testInput))
	require.Equal(t, 744, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 81, Part2(testInput))
	require.Equal(t, 1651, Part2(realInput))
}

var realInput = common.GetInput(10)

const testInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

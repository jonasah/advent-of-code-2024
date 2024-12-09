package day9

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 1928, Part1(testInput))
	require.Equal(t, 6320029754031, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 2858, Part2(testInput))
	require.Equal(t, 6347435485773, Part2(realInput))
}

var realInput = common.GetInput(9)

const testInput = `2333133121414131402`

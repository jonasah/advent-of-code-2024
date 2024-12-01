package day1

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 11, Part1(testInput))
	require.Equal(t, 2176849, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 31, Part2(testInput))
	require.Equal(t, 23384288, Part2(realInput))
}

var realInput = common.GetInput(1)

const testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

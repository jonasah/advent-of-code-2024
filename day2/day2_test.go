package day2

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 2, Part1(testInput))
	require.Equal(t, 564, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 4, Part2(testInput))
	require.Equal(t, 604, Part2(realInput))
}

var realInput = common.GetInput(2)

const testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

package day22

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 37327623, Part1(testInput))
	require.Equal(t, 13429191512, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 23, Part2(testInput2))
	require.Equal(t, 1582, Part2(realInput))
}

var realInput = common.GetInput(22)

const testInput = `1
10
100
2024`

const testInput2 = `1
2
3
2024`

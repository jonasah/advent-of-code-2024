package day19

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 6, Part1(testInput))
	require.Equal(t, 251, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 16, Part2(testInput))
	require.Equal(t, 616957151871345, Part2(realInput))
}

var realInput = common.GetInput(19)

const testInput = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

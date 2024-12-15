package day14

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 12, Part1(testInput, 11, 7))
	require.Equal(t, 217132650, Part1(realInput, 101, 103))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 6516, Part2(realInput, 101, 103))
}

var realInput = common.GetInput(14)

const testInput = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

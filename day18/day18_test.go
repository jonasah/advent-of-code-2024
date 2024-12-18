package day18

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 22, Part1(testInput, 6, 12))
	require.Equal(t, 286, Part1(realInput, 70, 1024))
}

func TestPart2(t *testing.T) {
	require.Equal(t, "6,1", Part2(testInput, 6))
	require.Equal(t, "20,64", Part2(realInput, 70))
}

var realInput = common.GetInput(18)

const testInput = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

package day3

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 161, Part1(testInput1))
	require.Equal(t, 190604937, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 48, Part2(testInput2))
	require.Equal(t, 82857512, Part2(realInput))
}

var realInput = common.GetInput(3)

const testInput1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const testInput2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

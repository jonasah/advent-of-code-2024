package day13

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 480, Part1(testInput))
	require.Equal(t, 36870, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 78101482023732, Part2(realInput))
}

var realInput = common.GetInput(13)

const testInput = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

package main

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, "4,6,3,5,6,3,5,2,1,0", Part1(testInput))
	require.Equal(t, "3,4,3,1,7,6,5,6,0", Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 117440, Part2(testInput2))
	require.Equal(t, 109019930331546, Part2(realInput))
}

var realInput = common.GetInput(17)

const testInput = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

const testInput2 = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

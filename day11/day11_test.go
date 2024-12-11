package day11

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 55312, Solve(testInput, 25))
	require.Equal(t, 207683, Solve(realInput, 25))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 244782991106220, Solve(realInput, 75))
}

var realInput = common.GetInput(11)

const testInput = `125 17`

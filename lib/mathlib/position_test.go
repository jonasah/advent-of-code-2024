package mathlib_test

import (
	"testing"

	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
	"github.com/stretchr/testify/assert"
)

func TestDirection_TurnRight(t *testing.T) {
	tests := map[mathlib.Direction]mathlib.Direction{
		mathlib.Right: mathlib.Down,
		mathlib.Down:  mathlib.Left,
		mathlib.Left:  mathlib.Up,
		mathlib.Up:    mathlib.Right,
	}

	for dir, expected := range tests {
		assert.Equal(t, expected, dir.TurnRight())
	}
}

func TestDirection_TurnLeft(t *testing.T) {
	tests := map[mathlib.Direction]mathlib.Direction{
		mathlib.Right: mathlib.Up,
		mathlib.Up:    mathlib.Left,
		mathlib.Left:  mathlib.Down,
		mathlib.Down:  mathlib.Right,
	}

	for dir, expected := range tests {
		assert.Equal(t, expected, dir.TurnLeft())
	}
}

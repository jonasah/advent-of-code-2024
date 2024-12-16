package mathlib

import (
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

func Sum[T constraints.Integer](a []T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

func ManhattanDistance(p0, p1 Position) int {
	return Abs(p0.X()-p1.X()) + Abs(p0.Y()-p1.Y())
}

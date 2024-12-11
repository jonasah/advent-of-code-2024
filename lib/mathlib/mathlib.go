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

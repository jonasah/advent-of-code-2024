package day11

import (
	"maps"
	"math"
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

func Solve(input string, blinks int) int {
	stones := slicelib.Atoi(strings.Fields(input))

	m := map[int]int{}
	for _, n := range stones {
		m[n]++
	}

	for range blinks {
		next := make(map[int]int, len(m))

		for n, count := range m {
			if n == 0 {
				next[1] += count
				continue
			}

			numDigits := int(math.Floor(math.Log10(float64(n)))) + 1

			if numDigits%2 == 0 {
				p := int(math.Pow10(numDigits / 2))
				left := n / p
				right := n % p
				next[left] += count
				next[right] += count
				continue
			}

			next[n*2024] += count
		}

		m = next
	}

	return mathlib.Sum(slices.Collect(maps.Values(m)))
}

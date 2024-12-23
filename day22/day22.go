package day22

import (
	"maps"
	"slices"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

type Sequence struct {
	a, b, c, d int
}

func Part1(input string) int {
	initial := slicelib.Atoi(common.GetLines(input))

	return mathlib.Sum(slicelib.Map(slicelib.Map(initial, calcNumbers), func(a []int) int { return a[len(a)-1] }))
}

func Part2(input string) int {
	initial := slicelib.Atoi(common.GetLines(input))

	bananas := make(map[Sequence]int)
	for _, n := range initial {
		prices := slicelib.Map(calcNumbers(n), func(n int) int { return n % 10 })

		diff := make([]int, 0, len(prices)-1)
		for i := range prices[1:] {
			diff = append(diff, prices[i+1]-prices[i])
		}

		handled := make(map[Sequence]bool, len(diff)-3)
		for i := range diff[3:] {
			seq := Sequence{diff[i], diff[i+1], diff[i+2], diff[i+3]}

			if !handled[seq] {
				handled[seq] = true
				bananas[seq] += prices[i+4]
			}
		}
	}

	return slices.Max(slices.Collect(maps.Values(bananas)))
}

func calcNumbers(initial int) []int {
	n := initial

	m := make([]int, 0, 2001)
	m = append(m, n)

	for range 2000 {
		n ^= n * 64
		n %= 16777216

		n ^= n / 32
		n %= 16777216

		n ^= n * 2048
		n %= 16777216

		m = append(m, n)
	}

	return m
}

package day18

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

func Part1(input string, maxCoord, numBytes int) int {
	return solve(parseInput(input)[:numBytes], maxCoord)
}

func Part2(input string, maxCoord int) string {
	positions := parseInput(input)

	left := 0
	half := len(positions) / 2

	for range 5000 {
		pivot := left + half
		d1 := solve(positions[:pivot], maxCoord)
		d2 := solve(positions[:pivot+1], maxCoord)

		if d1 != math.MaxInt && d2 == math.MaxInt {
			break
		}

		if d1 != math.MaxInt && d2 != math.MaxInt {
			left += half + 1
		}

		half = half / 2
	}

	p := positions[left+half]
	return fmt.Sprintf("%d,%d", p.X(), p.Y())
}

func parseInput(input string) []mathlib.Position {
	return slicelib.Map(
		common.GetLines(input),
		func(l string) mathlib.Position {
			xy := slicelib.Atoi(strings.Split(l, ","))
			return mathlib.NewPosition(xy[0], xy[1])
		})
}

func solve(bytes []mathlib.Position, maxCoord int) int {
	unvisited := make([]mathlib.Position, 0)
	dist := make(map[mathlib.Position]int)

	for x := range maxCoord + 1 {
		for y := range maxCoord + 1 {
			p := mathlib.NewPosition(x, y)

			if !slices.Contains(bytes, p) {
				unvisited = append(unvisited, p)
				dist[p] = math.MaxInt
			}
		}
	}

	end := mathlib.NewPosition(maxCoord, maxCoord)
	dist[mathlib.NewPosition(0, 0)] = 0

	dir := []mathlib.Direction{mathlib.Up, mathlib.Right, mathlib.Down, mathlib.Left}

	for len(unvisited) > 0 {
		curr := slices.MinFunc(unvisited, func(p0, p1 mathlib.Position) int { return dist[p0] - dist[p1] })

		if curr == end || dist[curr] == math.MaxInt {
			break
		}

		for _, d := range dir {
			n := curr.Move(d)
			if slices.Contains(unvisited, n) {
				dist[n] = mathlib.Min(dist[n], dist[curr]+1)
			}
		}

		unvisited = slices.DeleteFunc(unvisited, func(p mathlib.Position) bool { return p == curr })
	}

	return dist[end]
}

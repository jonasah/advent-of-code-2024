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

	b := make([]int, 0, len(positions))
	for i := range positions {
		b = append(b, i+1)
	}

	i, _ := slices.BinarySearchFunc(b, 0, func(numBytes, _ int) int {
		return solve(positions[:numBytes], maxCoord) - math.MaxInt
	})
	return fmt.Sprintf("%d,%d", positions[i].X(), positions[i].Y())
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

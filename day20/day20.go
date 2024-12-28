package day20

import (
	"slices"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

var dir = []mathlib.Direction{mathlib.Up, mathlib.Right, mathlib.Down, mathlib.Left}

type Candidate struct {
	p0, p1 mathlib.Position
}

func Part1(input string, threshold int) int {
	return solve(input, 2, threshold)
}

func Part2(input string, threshold int) int {
	return solve(input, 20, threshold)
}

func solve(input string, maxCheat, threshold int) int {
	positions, start := parseInput(input)
	track := createTrack(positions, start)

	return countCheats(track, maxCheat, threshold)
}

func parseInput(input string) ([]mathlib.Position, mathlib.Position) {
	lines := common.GetLines(input)

	var positions []mathlib.Position
	var start mathlib.Position

	for y, l := range lines {
		for x, t := range l {
			if t == '#' {
				continue
			}

			p := mathlib.NewPosition(x, y)
			positions = append(positions, p)

			if t == 'S' {
				start = p
			}
		}
	}

	return positions, start
}

func createTrack(positions []mathlib.Position, start mathlib.Position) []mathlib.Position {
	track := make([]mathlib.Position, 0, len(positions))
	curr := start

	for {
		track = append(track, curr)
		positions = slices.DeleteFunc(positions, func(p mathlib.Position) bool { return curr == p })

		var exists bool
		curr, exists = slicelib.FindFunc(positions, func(p mathlib.Position) bool { return mathlib.ManhattanDistance(curr, p) == 1 })
		if !exists {
			break
		}
	}

	return track
}

func countCheats(track []mathlib.Position, maxCheat, threshold int) int {
	count := 0
	for l0, p0 := range track {
		for l1, p1 := range track[l0+1:] {
			trackDist := l1 + 1
			shortestDist := mathlib.ManhattanDistance(p0, p1)
			saved := trackDist - shortestDist

			if shortestDist <= maxCheat && saved >= threshold {
				count++
			}
		}
	}

	return count
}

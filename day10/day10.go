package day10

import (
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

type Map struct {
	width, height int
	tiles         [][]int
	trailheads    []mathlib.Position
}

func (m *Map) heightAt(p mathlib.Position) int {
	return m.tiles[p.Y()][p.X()]
}

func Part1(input string) int {
	result, _ := solve(input)
	return result
}

func Part2(input string) int {
	_, result := solve(input)
	return result
}

func solve(input string) (int, int) {
	m := parseMap(input)
	topLeft := mathlib.NewPosition(0, 0)
	bottomRight := mathlib.NewPosition(m.width-1, m.height-1)

	dir := []mathlib.Direction{mathlib.Up, mathlib.Right, mathlib.Down, mathlib.Left}

	totalScore := 0
	totalRating := 0
	for _, t := range m.trailheads {
		trails := []mathlib.Position{t}
		trailEnds := map[mathlib.Position]any{}

		for len(trails) > 0 {
			p := trails[0]
			trails = trails[1:]

			if m.heightAt(p) == 9 {
				trailEnds[p] = true
				totalRating++
				continue
			}

			for _, d := range dir {
				next := p.Move(d)
				if next.IsWithin(topLeft, bottomRight) && m.heightAt(next) == m.heightAt(p)+1 {
					trails = append(trails, next)
				}
			}
		}

		totalScore += len(trailEnds)
	}

	return totalScore, totalRating
}

func parseMap(input string) Map {
	lines := common.GetLines(input)

	m := Map{width: len(lines[0]), height: len(lines)}
	for y, l := range lines {
		row := slicelib.Atoi(strings.Split(l, ""))
		m.tiles = append(m.tiles, row)

		trailheads := slicelib.IndexAll(row, 0)
		m.trailheads = append(m.trailheads, slicelib.Map(trailheads, func(x int) mathlib.Position { return mathlib.NewPosition(x, y) })...)
	}

	return m
}

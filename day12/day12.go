package day12

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

type Map struct {
	width, height int
	plants        [][]string
}

type Region struct {
	plant string
	plots []mathlib.Position
	perim int
}

func (m *Map) plantAt(p mathlib.Position) string {
	return m.plants[p.Y()][p.X()]
}

func Part1(input string) int {
	regions := getRegions(parseMap(input))

	cost := 0
	for _, r := range regions {
		cost += len(r.plots) * r.perim
	}

	return cost
}

func Part2(input string) int {
	plotMap := parseMap(input)

	newMap := Map{width: plotMap.width * 3, height: plotMap.height * 3}
	for _, row := range plotMap.plants {
		newRow := make([]string, 0, newMap.width)
		for _, plant := range row {
			newRow = append(newRow, plant, plant, plant)
		}

		newMap.plants = append(newMap.plants, newRow)
		newMap.plants = append(newMap.plants, newRow)
		newMap.plants = append(newMap.plants, newRow)
	}

	regions := getRegions(newMap)

	allDir := []mathlib.Direction{mathlib.Up, mathlib.Right, mathlib.Down, mathlib.Left}

	cost := 0
	for _, r := range regions {
		m := make(map[mathlib.Position]bool, len(r.plots))
		for _, p := range r.plots {
			m[p] = true
		}

		edgePositions := slicelib.FilterFunc(r.plots, func(p mathlib.Position) bool {
			return slices.ContainsFunc(allDir, func(d mathlib.Direction) bool { return !m[p.Move(d)] })
		})

		sides := 0

		for len(edgePositions) > 0 {
			p := slices.MinFunc(edgePositions, comparePos)

			up := !m[p.Move(mathlib.Up)]
			down := !m[p.Move(mathlib.Down)]

			d := mathlib.Down
			if up || down {
				d = mathlib.Right
			}

			for {
				edgePositions = slices.DeleteFunc(edgePositions, func(pp mathlib.Position) bool { return pp == p })

				p = p.Move(d)
				if !slices.Contains(edgePositions, p) {
					sides++
					break
				}
			}
		}

		cost += len(r.plots) / 9 * sides
	}

	return cost
}

func parseMap(input string) Map {
	lines := common.GetLines(input)

	m := Map{width: len(lines[0]), height: len(lines)}
	for _, l := range lines {
		m.plants = append(m.plants, strings.Split(l, ""))
	}

	return m
}

func getRegions(m Map) []Region {
	topLeft := mathlib.NewPosition(0, 0)
	bottomRight := mathlib.NewPosition(m.width-1, m.height-1)

	dir := []mathlib.Direction{mathlib.Up, mathlib.Right, mathlib.Down, mathlib.Left}
	visited := make(map[mathlib.Position]bool, m.width*m.height)

	var regions []Region
	for y := range m.height {
		for x := range m.width {
			p := mathlib.NewPosition(x, y)

			if visited[p] {
				continue
			}

			region := Region{plant: m.plantAt(p)}
			toVisit := []mathlib.Position{p}

			for len(toVisit) > 0 {
				curr := toVisit[0]
				toVisit = toVisit[1:]

				if visited[curr] {
					continue
				}

				visited[curr] = true
				region.plots = append(region.plots, curr)

				for _, d := range dir {
					next := curr.Move(d)
					if !next.IsWithin(topLeft, bottomRight) || m.plantAt(next) != region.plant {
						region.perim++
						continue
					}

					toVisit = append(toVisit, next)
				}
			}

			regions = append(regions, region)
		}
	}

	return regions
}

func comparePos(p0, p1 mathlib.Position) int {
	dx := p0.X() - p1.X()
	if dx != 0 {
		return dx
	}

	return p0.Y() - p1.Y()
}

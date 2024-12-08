package day8

import (
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
)

type Map struct {
	width, height int
	antennas      map[string][]mathlib.Position
}

func Part1(input string) int {
	m := parseMap(input)

	topLeft := mathlib.NewPosition(0, 0)
	bottomRight := mathlib.NewPosition(m.width-1, m.height-1)

	antinodes := map[mathlib.Position]any{}
	for _, positions := range m.antennas {
		for i, p0 := range positions {
			for _, p1 := range positions[i+1:] {
				d := p1.Diff(p0)

				a0 := p0.Move(-d)
				if a0.IsWithin(topLeft, bottomRight) {
					antinodes[a0] = true
				}

				a1 := p1.Move(d)
				if a1.IsWithin(topLeft, bottomRight) {
					antinodes[a1] = true
				}
			}
		}
	}

	return len(antinodes)
}

func Part2(input string) int {
	m := parseMap(input)

	topLeft := mathlib.NewPosition(0, 0)
	bottomRight := mathlib.NewPosition(m.width-1, m.height-1)

	antinodes := map[mathlib.Position]any{}
	for _, positions := range m.antennas {
		for i, p0 := range positions {
			for _, p1 := range positions[i+1:] {
				d := p1.Diff(p0)

				a0 := p0
				for a0.IsWithin(topLeft, bottomRight) {
					antinodes[a0] = true
					a0 = a0.Move(-d)
				}

				a1 := p1
				for a1.IsWithin(topLeft, bottomRight) {
					antinodes[a1] = true
					a1 = a1.Move(d)
				}
			}
		}
	}

	return len(antinodes)
}

func parseMap(input string) Map {
	lines := common.GetLines(input)

	m := Map{
		width:    len(lines[0]),
		height:   len(lines),
		antennas: make(map[string][]mathlib.Position),
	}
	for r, l := range lines {
		for c, v := range strings.Split(l, "") {
			if v != "." {
				m.antennas[v] = append(m.antennas[v], mathlib.NewPosition(c, r))
			}
		}
	}

	return m
}

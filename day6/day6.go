package day6

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/common"
	"github.com/jonasah/advent-of-code-2024/lib/mathlib"
)

const (
	guard    = "^"
	obstacle = "#"
)

type Map struct {
	width, height int
	obstacles     map[mathlib.Position]any
	empty         []mathlib.Position
}

func Part1(input string) int {
	m, guardPos := parseMap(input)
	guardDir := mathlib.Up

	topLeft := mathlib.NewPosition(0, 0)
	bottomRight := mathlib.NewPosition(m.width-1, m.height-1)

	visited := map[mathlib.Position]any{}

	for guardPos.IsWithin(topLeft, bottomRight) {
		visited[guardPos] = true

		_, obstacleAhead := m.obstacles[guardPos.Move(guardDir)]
		if obstacleAhead {
			guardDir = guardDir.TurnRight()
		}

		guardPos = guardPos.Move(guardDir)
	}

	return len(visited)
}

func Part2(input string) int {
	m, origGuardPos := parseMap(input)
	origGuardDir := mathlib.Up

	topLeft := mathlib.NewPosition(0, 0)
	bottomRight := mathlib.NewPosition(m.width-1, m.height-1)

	result := 0

	for _, e := range m.empty {
		m.obstacles[e] = true

		guardPos := origGuardPos
		guardDir := origGuardDir

		visited := map[mathlib.Position][]mathlib.Direction{}

		for guardPos.IsWithin(topLeft, bottomRight) {
			if slices.Contains(visited[guardPos], guardDir) {
				result++
				break
			}

			visited[guardPos] = append(visited[guardPos], guardDir)

			_, obstacleAhead := m.obstacles[guardPos.Move(guardDir)]
			if obstacleAhead {
				guardDir = guardDir.TurnRight()
			}

			_, obstacleAhead = m.obstacles[guardPos.Move(guardDir)]
			if !obstacleAhead {
				guardPos = guardPos.Move(guardDir)
			}
		}

		delete(m.obstacles, e)
	}

	return result
}

func parseMap(input string) (Map, mathlib.Position) {
	lines := common.GetLines(input)

	m := Map{
		width:     len(lines[0]),
		height:    len(lines),
		obstacles: make(map[mathlib.Position]any),
	}
	var guardPos mathlib.Position
	for r, l := range lines {
		for c, v := range strings.Split(l, "") {
			p := mathlib.NewPosition(c, r)

			switch v {
			case guard:
				guardPos = p
			case obstacle:
				m.obstacles[p] = true
			default:
				m.empty = append(m.empty, p)
			}
		}
	}

	return m, guardPos
}
